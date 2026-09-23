package repository

import (
	"errors"

	"github.com/sangeeth518/go-Ecommerce/pkg/domain"
	interfaces "github.com/sangeeth518/go-Ecommerce/pkg/repository/interface"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
	"gorm.io/gorm"
)

type orderRepo struct {
	DB *gorm.DB
}

func NewOrderRepo(db *gorm.DB) interfaces.OrderRepo {
	return &orderRepo{
		DB: db,
	}
}

func (o *orderRepo) GetAddressByID(addressID, userID int) (domain.Adress, error) {
	var adress domain.Adress

	err := o.DB.Where("user_id =? AND id=?", userID, addressID).First(&adress).Error
	if err != nil {
		return domain.Adress{}, err
	}

	return adress, nil
}

func (o *orderRepo) GetCartByUserID(userID int) (domain.Cart, error) {
	var cart domain.Cart

	err := o.DB.Where("user_id =?", userID).First(&cart).Error
	if err != nil {
		return domain.Cart{}, err
	}

	return cart, nil
}

func (o *orderRepo) GetCartItems(cartID int) ([]models.CartItemResponse, error) {
	var cartItems []models.CartItemResponse

	query := `SELECT ci.inventory_id,i.product_name, i.size, ci.quantity,i.price , ci.total_price From cart_items ci JOIN inventories i on ci.inventory_id =i.id where ci.cart_id =? `

	err := o.DB.Raw(query, cartID).Scan(&cartItems).Error
	if err != nil {
		return []models.CartItemResponse{}, err
	}

	return cartItems, nil
}

func (o *orderRepo) CreateCODOrder(order domain.Order, items []domain.OrderItem) (domain.Order, error) {
	tx := o.DB.Begin()
	if tx.Error != nil {
		return domain.Order{}, tx.Error
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return domain.Order{}, err
	}

	// Create order items
	for _, item := range items {
		item.OrderId = order.Id
		if err := tx.Create(&item).Error; err != nil {
			tx.Rollback()
			return domain.Order{}, err
		}

		// Reduce stock
		result := tx.Model(&domain.Inventory{}).Where("id =? AND stock >=?", item.InventoryId, item.Quantity).Update("stock", gorm.Expr("stock -?", item.Quantity))
		if result.Error != nil {
			tx.Rollback()
			return domain.Order{}, result.Error
		}
		if result.RowsAffected == 0 {
			tx.Rollback()
			return domain.Order{}, errors.New("insufficient stock for inventory")
		}
	}

	// Clear cart items
	query := `DELETE FROM cart_items WHERE cart_id = (SELECT id FROM carts WHERE user_id =?)`
	if err := tx.Exec(query, order.UserId).Error; err != nil {
		tx.Rollback()
		return domain.Order{}, err
	}

	if err := tx.Model(&domain.Cart{}).Where("user_id = ?", order.UserId).Update("total_price", 0).Error; err != nil {
		tx.Rollback()
		return domain.Order{}, err
	}

	if err := tx.Commit().Error; err != nil {
		return domain.Order{}, err
	}

	return order, nil
}

func (o *orderRepo) CreatePendingOrder(order domain.Order, items []domain.OrderItem) (domain.Order, error) {
	err := o.DB.Transaction(func(tx *gorm.DB) error {
		// 1. Create order record
		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		// 2. Insert order items linked to the new order ID
		for _, item := range items {
			item.OrderId = order.Id
			if err := tx.Create(&item).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}

func (o *orderRepo) ConfirmPaymentAndReduceStock(razorpayOrderID, paymentID string) error {
	return o.DB.Transaction(func(tx *gorm.DB) error {
		// 1. Find the pending order by razorpay_order_id
		var order domain.Order
		if err := tx.Where("razorpay_order_id = ?", razorpayOrderID).First(&order).Error; err != nil {
			return errors.New("order not found for this razorpay order ID")
		}

		// 2. Fetch all items belonging to this order
		var items []domain.OrderItem
		if err := tx.Where("order_id = ?", order.Id).Find(&items).Error; err != nil {
			return err
		}

		// 3. Atomically reduce inventory stock for each item
		for _, item := range items {
			res := tx.Model(&domain.Inventory{}).
				Where("id = ? AND stock >= ?", item.InventoryId, item.Quantity).
				Update("stock", gorm.Expr("stock - ?", item.Quantity))

			if res.Error != nil {
				return res.Error
			}
			if res.RowsAffected == 0 {
				return errors.New("insufficient stock to confirm payment")
			}
		}

		// 4. Update order status to confirmed and payment to paid
		if err := tx.Model(&domain.Order{}).Where("id = ?", order.Id).Updates(map[string]interface{}{
			"order_status":        "confirmed",
			"payment_status":      "paid",
			"razorpay_payment_id": paymentID,
		}).Error; err != nil {
			return err
		}

		// 5. Clear the user's cart items
		query := `DELETE FROM cart_items WHERE cart_id = (SELECT id FROM carts WHERE user_id = ?)`
		if err := tx.Exec(query, order.UserId).Error; err != nil {
			return err
		}

		// 6. Reset cart total price
		if err := tx.Model(&domain.Cart{}).Where("user_id = ?", order.UserId).Update("total_price", 0).Error; err != nil {
			return err
		}

		return nil
	})
}

func (o *orderRepo) UpdatePaymentFailure(razorpayOrderID string) error {
	return o.DB.Model(&domain.Order{}).
		Where("razorpay_order_id = ?", razorpayOrderID).
		Updates(map[string]interface{}{
			"order_status":   "cancelled",
			"payment_status": "failed",
		}).Error
}

func (o *orderRepo) GetOrdersByUserID(userID int) ([]models.MyOrdersResponse, error) {
	var orders []models.MyOrdersResponse
	query := `
		SELECT 
			o.id,
			o.final_price,
			o.payment_method,
			o.payment_status,
			o.order_status,
			o.created_at,
			COUNT(oi.id) AS item_count
		FROM orders o
		LEFT JOIN order_items oi ON o.id = oi.order_id
		WHERE o.user_id = ?
		GROUP BY o.id
		ORDER BY o.created_at DESC
	`
	err := o.DB.Raw(query, userID).Scan(&orders).Error
	if err != nil {
		return []models.MyOrdersResponse{}, err
	}
	return orders, nil
}

func (o *orderRepo) GetOrderDetails(orderID, userID int) (models.OrderDetailsResponse, error) {
	var order models.OrderResponse
	orderQuery := `
		SELECT id, user_id, address_id, payment_method, payment_status, order_status, final_price, discount, razorpay_order_id, created_at
		FROM orders
		WHERE id = ? AND user_id = ?
	`
	if err := o.DB.Raw(orderQuery, orderID, userID).Scan(&order).Error; err != nil {
		return models.OrderDetailsResponse{}, err
	}
	if order.Id == 0 {
		return models.OrderDetailsResponse{}, errors.New("order not found")
	}

	// Fetch Shipping Address
	var address models.AddressResponse
	addressQuery := `
		SELECT name, house_name, street, city, state, phone, pin
		FROM adresses
		WHERE id = ?
	`
	if err := o.DB.Raw(addressQuery, order.AddressId).Scan(&address).Error; err != nil {
		return models.OrderDetailsResponse{}, err
	}

	// Fetch Order Items with Product details
	var items []models.OrderItemResponse
	itemsQuery := `
		SELECT 
			oi.inventory_id,
			i.product_name,
			i.size,
			oi.quantity,
			oi.unit_price,
			oi.total_price
		FROM order_items oi
		JOIN inventories i ON oi.inventory_id = i.id
		WHERE oi.order_id = ?
	`
	if err := o.DB.Raw(itemsQuery, order.Id).Scan(&items).Error; err != nil {
		return models.OrderDetailsResponse{}, err
	}

	return models.OrderDetailsResponse{
		Order:      order,
		Address:    address,
		OrderItems: items,
	}, nil
}

// CancelOrder cancels an active order and adds the ordered quantities back to inventory stock
func (o *orderRepo) CancelOrder(orderID, userID int) error {
	return o.DB.Transaction(func(tx *gorm.DB) error {
		var order domain.Order
		if err := tx.Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
			return errors.New("order not found")
		}

		if order.OrderStatus == "cancelled" {
			return errors.New("order is already cancelled")
		}
		if order.OrderStatus == "delivered" {
			return errors.New("cannot cancel an order that has already been delivered")
		}

		// Fetch items for this order to restock
		var items []domain.OrderItem
		if err := tx.Where("order_id = ?", order.Id).Find(&items).Error; err != nil {
			return err
		}

		// Restock each inventory item (stock = stock + quantity)
		for _, item := range items {
			if err := tx.Model(&domain.Inventory{}).
				Where("id = ?", item.InventoryId).
				Update("stock", gorm.Expr("stock + ?", item.Quantity)).Error; err != nil {
				return err
			}
		}

		// Update order status to cancelled
		if err := tx.Model(&domain.Order{}).Where("id = ?", order.Id).Updates(map[string]interface{}{
			"order_status":   "cancelled",
			"payment_status": "cancelled",
		}).Error; err != nil {
			return err
		}

		return nil
	})
}

func (o *orderRepo) GetUserDetails(userID int) (models.UserDetailsResponse, error) {
	var user models.UserDetailsResponse
	query := `SELECT id, name, email, phone FROM users WHERE id = ?`
	err := o.DB.Raw(query, userID).Scan(&user).Error
	if err != nil {
		return models.UserDetailsResponse{}, err
	}
	return user, nil
}
