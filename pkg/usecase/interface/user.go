package interfaces

import "github.com/sangeeth518/go-Ecommerce/pkg/utils/models"

type UserUsecase interface {
	UserSignup(user models.UserDetails) (models.UserToken, error)
}
