package interfaces

import "github.com/sangeeth518/go-Ecommerce/pkg/utils/models"

type UserRepo interface {
	UserSignup(user models.UserDetails) (models.UserDetailsResponse, error)
	CheckUserAvailability(email string) bool
	UserBlockStatus(emil string) (bool, error)
	FindUserByEmail(user models.UserLogin) (models.UserSigninResponse, error)
	Changepassword(id int, passowrd string) error
	GetPassword(id int) (string, error)
	AddAdress(id int, adress models.AddAdress) error
}
