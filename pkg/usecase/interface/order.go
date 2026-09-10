package interfaces

import "github.com/sangeeth518/go-Ecommerce/pkg/utils/models"

type OrderUsecase interface {
	OrderCheckout(userID int, input models.OrderIncoming) (models.OrderResponse, error)
	VerifyPayment(payment models.PaymentVerification) error
	GetMyOrders(userID int) ([]models.MyOrdersResponse, error)
	GetOrderDetails(orderID, userID int) (models.OrderDetailsResponse, error)
}
