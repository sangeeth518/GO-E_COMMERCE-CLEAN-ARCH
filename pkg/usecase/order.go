package usecase

import (
	"errors"
	"strings"

	"github.com/sangeeth518/go-Ecommerce/pkg/domain"
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

	_, err := o.orderrepo.GetAddressByID(input.AddressId, userID)
	if err != nil {
		return models.OrderResponse{}, errors.New("invalid address")
	}

	cart, err := o.orderrepo.GetCartByUserID(userID)
	if err != nil {
		return models.OrderResponse{}, errors.New("cart not found of user")

	}

	if cart.TotalPrice <= 0 {
		return models.OrderResponse{}, errors.New("cart is empty")
	}

	cartItems, err := o.orderrepo.GetCartItems(cart.Id)
	if err != nil || len(cartItems) == 0 {
		return models.OrderResponse{}, errors.New("no items found in cart")
	}

	var orderItems []domain.OrderItem

	for _, ci := range cartItems {
		orderItems = append(orderItems, domain.OrderItem{
			Quantity:    ci.Quantity,
			InventoryId: ci.InventoryId,
			UnitPrice:   ci.Price,
			TotalPrice:  ci.TotalPrice,
		})
	}

	paymentMethod := strings.ToUpper(strings.TrimSpace(input.PaymentMethod))
	if paymentMethod != "COD" {
		return models.OrderResponse{}, errors.New("only COD availale for now")
	}

	order := domain.Order{
		UserId:        userID,
		AddressId:     input.AddressId,
		PaymentMethod: "COD",
		PaymentStatus: "pending",
		OrderStatus:   "confirmed",
		FinalPrice:    cart.TotalPrice,
		Discount:      0,
	}

	placeOrder, err := o.orderrepo.CreateCODOrder(order, orderItems)
	if err != nil {
		return models.OrderResponse{}, err
	}

	return models.OrderResponse{
		Id:            placeOrder.Id,
		UserId:        placeOrder.UserId,
		AddressId:     placeOrder.AddressId,
		PaymentStatus: placeOrder.PaymentStatus,
		PaymentMethod: placeOrder.PaymentMethod,
		OrderStatus:   placeOrder.OrderStatus,
		FinalPrice:    placeOrder.FinalPrice,
		Discount:      placeOrder.Discount,
		CreatedAt:     placeOrder.CreatedAt,
	}, nil

}

func (o *orderUsecase) VerifyPayment(payment models.PaymentVerification) error {
	return nil
}

func (o *orderUsecase) GetMyOrders(userID int) ([]models.MyOrdersResponse, error) {
	return o.orderrepo.GetOrdersByUserID(userID)
}

func (o *orderUsecase) GetOrderDetails(orderID, userID int) (models.OrderDetailsResponse, error) {
	return o.orderrepo.GetOrderDetails(orderID, userID)
}

func (o *orderUsecase) CancelOrder(orderID, userID int) error {
	return o.orderrepo.CancelOrder(orderID, userID)
}

func (o *orderUsecase) PrintInvoice(orderID, userID int) ([]byte, error) {
	details, err := o.orderrepo.GetOrderDetails(orderID, userID)
	if err != nil {
		return nil, err
	}
	return o.helper.GenerateInvoicePDF(details)
}
