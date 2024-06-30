package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sangeeth518/go-Ecommerce/pkg/domain"
	interfaces "github.com/sangeeth518/go-Ecommerce/pkg/usecase/interface"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/response"
)

type CategoryHandler struct {
	category interfaces.CateoryUsecase
}

func NewCategoryHandler(c interfaces.CateoryUsecase) *CategoryHandler {
	return &CategoryHandler{
		category: c,
	}
}

func (ch *CategoryHandler) AddCategory(c *gin.Context) {
	var category domain.Category
	if err := c.BindJSON(&category); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "data not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	categoryresponse, err := ch.category.AddCategory(category)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "cannot add category ", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "category added successfully", categoryresponse, nil)
	c.JSON(http.StatusOK, successRes)

}
