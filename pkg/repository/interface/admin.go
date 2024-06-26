package interfaces

import (
	"github.com/sangeeth518/go-Ecommerce/pkg/domain"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
)

type AdminRepo interface {
	LoginHandler(admindetails models.AdminLogin) (domain.Admin, error)
	GetUserById(id string) (domain.User, error)
	BlockUserById(user domain.User) error
	GetUsers(page int, count int) ([]models.UserdetailsAtAdmin, error)
}
