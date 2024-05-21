package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	interfaces "github.com/sangeeth518/go-Ecommerce/pkg/usecase/interface"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/response"
)

type AdminHandler struct {
	adminUsecase interfaces.AdminUseCase
}

func NewAdminHandler(usecase interfaces.AdminUseCase) *AdminHandler {
	return &AdminHandler{
		adminUsecase: usecase,
	}
}

func (ad *AdminHandler) LoginHandler(c *gin.Context) {
	var admindetails models.AdminLogin
	if err := c.BindJSON(&admindetails); err != nil {
		errres := response.ClientResponse(http.StatusBadRequest, "details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errres)
	}

}
