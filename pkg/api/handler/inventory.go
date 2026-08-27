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

type InventoryHandler struct {
	inv interfaces.InventoryUsecase
}

func NewInventoryHandler(inv interfaces.InventoryUsecase) *InventoryHandler {
	return &InventoryHandler{
		inv: inv,
	}
}

func (i *InventoryHandler) AddProduct(c *gin.Context) {

	catIDstr := c.Param("category_id")
	catId, err := strconv.Atoi(catIDstr)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "wrong format of category id", nil, nil)
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	var product models.AddProduct
	if err := c.BindJSON(&product); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	product.CategoryId = catId

	if err := validator.New().Struct(product); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constraints not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	productResponse, err := i.inv.AddProduct(product)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to add product", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusCreated, "product added succesfully", productResponse, nil)
	c.JSON(http.StatusCreated, successRes)

}

func (i *InventoryHandler) ListProducts(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "limit number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	products, err := i.inv.ListProducts(page, limit)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve products", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully retrieved all products", products, nil)
	c.JSON(http.StatusOK, successRes)
}
