package interfaces

import "github.com/sangeeth518/go-Ecommerce/pkg/utils/models"

type Helper interface {
	GenerateTokenAdmin(admin models.AdminDetailResponse) (string, string, error)
	GenerateTokenClient(user models.UserDetailsResponse) (string, error)
	PasswordHashing(password string) (string, error)
	CompareHashPassword(password string, givenpass string) error
}
