package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	interfaces "github.com/sangeeth518/go-Ecommerce/pkg/usecase/interface"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/response"
)

type CartHandler struct {
	cartUsecase interfaces.CartUsecase
}

func NewCartHandler(usecase interfaces.CartUsecase) *CartHandler {
	return &CartHandler{cartUsecase: usecase}
}

func (h *CartHandler) AddToCart(c *gin.Context) {
	userID, exists := c.Get("id")
	if !exists {
		errRes := response.ClientResponse(http.StatusUnauthorized, "Unauthorized", nil, "user not authenticated")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}

	var req models.AddToCart
	if err := c.BindJSON(&req); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Invalid input format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := validator.New().Struct(req); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Validation error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := h.cartUsecase.AddToCart(userID.(int), req); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Could not add to cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added to cart", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (h *CartHandler) ViewCart(c *gin.Context) {
	userID, exists := c.Get("id")
	if !exists {
		errRes := response.ClientResponse(http.StatusUnauthorized, "Unauthorized", nil, "user not authenticated")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}

	cart, err := h.cartUsecase.ViewCart(userID.(int))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Could not retrieve cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully retrieved cart", cart, nil)
	c.JSON(http.StatusOK, successRes)
}

func (h *CartHandler) UpdateQuantity(c *gin.Context) {
	userID, exists := c.Get("id")
	if !exists {
		errRes := response.ClientResponse(http.StatusUnauthorized, "Unauthorized", nil, "user not authenticated")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}

	var req models.UpdateQuantityReq
	if err := c.BindJSON(&req); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Invalid input format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := validator.New().Struct(req); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Validation error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := h.cartUsecase.UpdateQuantity(userID.(int), req); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Could not update quantity", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully updated cart quantity", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (h *CartHandler) RemoveProductFromCart(c *gin.Context) {
	userID, exists := c.Get("id")
	if !exists {
		errRes := response.ClientResponse(http.StatusUnauthorized, "Unauthorized", nil, "user not authenticated")
		c.JSON(http.StatusUnauthorized, errRes)
		return
	}

	inventoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Invalid product id in URL", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := h.cartUsecase.RemoveProductFromCart(userID.(int), inventoryID); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "Could not remove item from cart", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully removed item from cart", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
