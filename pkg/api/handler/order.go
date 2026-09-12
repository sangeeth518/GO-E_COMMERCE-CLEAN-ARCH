package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	interfaces "github.com/sangeeth518/go-Ecommerce/pkg/usecase/interface"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/response"
)

type OrderHandler struct {
	orderUsecase interfaces.OrderUsecase
}

func NewOrderHandler(orderUsecase interfaces.OrderUsecase) *OrderHandler {
	return &OrderHandler{
		orderUsecase: orderUsecase,
	}
}

func (o *OrderHandler) OrderCheckout(c *gin.Context) {
	userId, ok := c.Get("id")
	if !ok {
		errRes := response.ClientResponse(http.StatusUnauthorized, "Unauthorized", nil, "user not authenticated")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}

	var input models.OrderIncoming
	if err := c.BindJSON(&input); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Invalid input format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := validator.New().Struct(input); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Validation error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	order, err := o.orderUsecase.OrderCheckout(userId.(int), input)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to place order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "Order placed successfully", order, nil)
	c.JSON(http.StatusCreated, successRes)
}

func (o *OrderHandler) GetMyOrder(c *gin.Context) {
	userId, ok := c.Get("id")
	if !ok {
		errRes := response.ClientResponse(http.StatusUnauthorized, "Unauthorized", nil, "user not authenticated")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}

	orders, err := o.orderUsecase.GetMyOrders(userId.(int))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to get my orders", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully retrieved my orders", orders, nil)
	c.JSON(http.StatusOK, successRes)

}

func (o *OrderHandler) GetOrderDetails(c *gin.Context) {

	userId, ok := c.Get("id")
	if !ok {
		errRes := response.ClientResponse(http.StatusUnauthorized, "Unauthorized", nil, "user not authenticated")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}

	orderId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Invalid order id in URL", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	details, err := o.orderUsecase.GetOrderDetails(orderId, userId.(int))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to get order details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully retrieved order details", details, nil)
	c.JSON(http.StatusOK, successRes)

}

func (o *OrderHandler) CancelOrder(c *gin.Context) {
	userId, ok := c.Get("id")
	if !ok {
		errRes := response.ClientResponse(http.StatusUnauthorized, "Unauthorized", nil, "user not authenticated")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}

	orderId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Invalid order id in URL", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := o.orderUsecase.CancelOrder(orderId, userId.(int)); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Failed to cancel order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Order cancelled successfully and stock restored", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (o *OrderHandler) DownloadInvoice(c *gin.Context) {
	userId, ok := c.Get("id")
	if !ok {
		errRes := response.ClientResponse(http.StatusUnauthorized, "Unauthorized", nil, "user not authenticated")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}

	orderId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Invalid order id in URL", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	pdfBytes, err := o.orderUsecase.PrintInvoice(orderId, userId.(int))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Failed to generate invoice", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=invoice_order_%d.pdf", orderId))
	c.Header("Content-Type", "application/pdf")
	c.Data(http.StatusOK, "application/pdf", pdfBytes)
}
