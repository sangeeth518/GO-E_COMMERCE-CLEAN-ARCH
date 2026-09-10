package usecase

import (
	interfaces "github.com/sangeeth518/go-Ecommerce/pkg/helper/interface"
	interfacee "github.com/sangeeth518/go-Ecommerce/pkg/repository/interface"
	services "github.com/sangeeth518/go-Ecommerce/pkg/usecase/interface"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
)

type orderUsecase struct {
	orderrepo interfacee.OrderRepo
	helper    interfaces.Helper
}

func NewOrderUsecase(orderrepo interfacee.OrderRepo, helper interfaces.Helper) services.OrderUsecase {
	return &orderUsecase{
		orderrepo: orderrepo,
		helper:    helper,
	}
}

func (o *orderUsecase) OrderCheckout(userID int, input models.OrderIncoming) (models.OrderResponse, error) {

	return models.OrderResponse{}, nil

}

func (o *orderUsecase) VerifyPayment(payment models.PaymentVerification) error {
	return nil
}

func (o *orderUsecase) GetMyOrders(userID int) ([]models.MyOrdersResponse, error) {
	return []models.MyOrdersResponse{}, nil
}

func (o *orderUsecase) GetOrderDetails(orderID, userID int) (models.OrderDetailsResponse, error) {
	return models.OrderDetailsResponse{}, nil
}
