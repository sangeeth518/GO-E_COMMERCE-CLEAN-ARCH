package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	interfaces "github.com/sangeeth518/go-Ecommerce/pkg/usecase/interface"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/response"
)

type UserHandler struct {
	userUsecase interfaces.UserUsecase
}

func NewUserHandler(user interfaces.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: user,
	}
}

func (u *UserHandler) UserSignup(c *gin.Context) {
	var user models.UserDetails
	if err := c.BindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
	}

	err := validator.New().Struct(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constraints not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
	}

	usercreated, err := u.userUsecase.UserSignup(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "user could't signup", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
	}

	successRes := response.ClientResponse(http.StatusCreated, "User succesfully signedup", usercreated, nil)
	c.JSON(http.StatusCreated, successRes)

}
