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

func (ch *CategoryHandler) ShowCategories(c *gin.Context) {
	cat, err := ch.category.ShowCategories()
	if err != nil {
		errres := response.ClientResponse(http.StatusBadRequest, "cant fetch categories", nil, err.Error())
		c.JSON(http.StatusBadRequest, errres)
	}
	succesres := response.ClientResponse(http.StatusOK, "Categories fetched succesfully", cat, nil)
	c.JSON(http.StatusOK, succesres)
}

// Edit catgegory(can only be done by admin)
func (ch *CategoryHandler) EditCategory(c *gin.Context) {
	var category domain.Category
	err := c.BindJSON(&category)
	if err != nil {
		errres := response.ClientResponse(http.StatusBadRequest, "Data not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errres)
		return
	}
	err = ch.category.EditCategory(category)
	if err != nil {
		errres := response.ClientResponse(http.StatusBadRequest, "Cannot update category", nil, err.Error())
		c.JSON(http.StatusBadRequest, errres)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Category Updated succesfully", nil, nil)
	c.JSON(http.StatusOK, successRes)

}
