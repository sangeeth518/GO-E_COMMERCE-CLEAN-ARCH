package repository

import (
	"errors"

	"github.com/sangeeth518/go-Ecommerce/pkg/domain"
	interfaces "github.com/sangeeth518/go-Ecommerce/pkg/repository/interface"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
	"gorm.io/gorm"
)

type CartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) interfaces.CartRepository {
	return &CartRepository{DB: db}
}

func (c *CartRepository) GetCartByUserID(userID int) (domain.Cart, error) {
	var cart domain.Cart
	err := c.DB.Where("user_id = ?", userID).First(&cart).Error
	return cart, err
}

func (c *CartRepository) CreateCart(userID int) (domain.Cart, error) {
	cart := domain.Cart{UserId: userID, TotalPrice: 0}
	err := c.DB.Create(&cart).Error
	return cart, err
}

func (c *CartRepository) GetCartItem(cartID, inventoryID int) (domain.CartItem, error) {
	var item domain.CartItem
	err := c.DB.Where("cart_id = ? AND inventory_id = ?", cartID, inventoryID).First(&item).Error
	return item, err
}

func (c *CartRepository) AddCartItem(item domain.CartItem) error {
	return c.DB.Create(&item).Error
}

func (c *CartRepository) UpdateCartItem(item domain.CartItem) error {
	return c.DB.Model(&domain.CartItem{}).
		Where("cart_id = ? AND inventory_id = ?", item.CartId, item.InventoryId).
		Updates(map[string]interface{}{
			"quantity":    item.Quantity,
			"total_price": item.TotalPrice,
		}).Error
}

func (c *CartRepository) RemoveCartItem(cartID, inventoryID int) error {
	result := c.DB.Where("cart_id = ? AND inventory_id = ?", cartID, inventoryID).Delete(&domain.CartItem{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("item not found in cart")
	}
	return nil
}

func (c *CartRepository) GetCartItems(cartID int) ([]models.CartItemResponse, error) {
	var items []models.CartItemResponse
	query := `
		SELECT 
			ci.inventory_id,
			i.product_name,
			i.size,
			ci.quantity,
			i.price,
			ci.total_price
		FROM cart_items ci
		JOIN inventories i ON ci.inventory_id = i.id
		WHERE ci.cart_id = ?
	`
	err := c.DB.Raw(query, cartID).Scan(&items).Error
	return items, err
}

func (c *CartRepository) GetProductStockAndPrice(inventoryID int) (int, float64, error) {
	var result struct {
		Stock int
		Price float64
	}
	err := c.DB.Model(&domain.Inventory{}).
		Select("stock, price").
		Where("id = ?", inventoryID).
		Scan(&result).Error
	return result.Stock, result.Price, err
}

func (c *CartRepository) UpdateCartTotal(cartID int, total float64) error {
	return c.DB.Model(&domain.Cart{}).Where("id = ?", cartID).Update("total_price", total).Error
}
