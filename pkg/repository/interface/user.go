package interfaces

import "github.com/sangeeth518/go-Ecommerce/pkg/utils/models"

type UserRepo interface {
	UserSignup(user models.UserDetails) (models.UserDetailsResponse, error)
	CheckUserAvailability(email string) bool
}
