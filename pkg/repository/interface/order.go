package interfaces

import (
	"github.com/sangeeth518/go-Ecommerce/pkg/domain"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
)

type OrderRepo interface {
	// Address & Cart validation checks
	GetAddressByID(addressID, userID int) (domain.Adress, error)
	GetCartByUserID(userID int) (domain.Cart, error)
	GetCartItems(cartID int) ([]models.CartItemResponse, error)

	// Checkout transactions
	CreateCODOrder(order domain.Order, items []domain.OrderItem) (domain.Order, error)
	CreatePendingOrder(order domain.Order, items []domain.OrderItem) (domain.Order, error)

	// Payment verification & stock finalization for Razorpay
	ConfirmPaymentAndReduceStock(razorpayOrderID, paymentID string) error
	UpdatePaymentFailure(razorpayOrderID string) error

	// Order details & history
	GetOrdersByUserID(userID int) ([]models.MyOrdersResponse, error)
	GetOrderDetails(orderID, userID int) (models.OrderDetailsResponse, error)
}
