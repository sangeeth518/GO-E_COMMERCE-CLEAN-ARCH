package handler

import (
	"net/http"
	"strconv"

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

func (ad *AdminHandler) BlockUser(c *gin.Context) {
	id := c.Param("id")
	err := ad.adminUsecase.BlockUser(id)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "user could not be blocked", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "succesfully blocked the user", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (ad *AdminHandler) UnblockUser(c *gin.Context) {
	id := c.Param("id")
	err := ad.adminUsecase.UnblockUser(id)

	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "user could not be unblocked", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "succesfully unblocked the user", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (a *AdminHandler) Getusers(c *gin.Context) {
	pagestr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pagestr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check path parameter", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return

	}
	count, err := strconv.Atoi(c.DefaultQuery("count", "2"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "check path parameter", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	users, err := a.adminUsecase.GetUsers(page, count)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusInternalServerError, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully retrieved the users", users, nil)
	c.JSON(http.StatusOK, successRes)

}
