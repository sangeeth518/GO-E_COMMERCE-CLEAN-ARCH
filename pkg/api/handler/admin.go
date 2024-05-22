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
		errRes := response.ClientResponse(http.StatusBadRequest, "details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	admin, err := ad.adminUsecase.LoginHandler(admindetails)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "cannot authenticate admin", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	c.SetCookie("Authorization", admin.Token, 3600*24*30, "", "", false, true)
	c.SetCookie("Refresh", admin.Refresh, 3600*24*30, "", "", false, true)

	successRes := response.ClientResponse(http.StatusOK, "Admin logged in succesfully", admin, nil)
	c.JSON(http.StatusOK, successRes)
}
