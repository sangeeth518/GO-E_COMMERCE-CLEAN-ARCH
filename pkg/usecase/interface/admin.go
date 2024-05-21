package interfaces

import (
	"github.com/sangeeth518/go-Ecommerce/pkg/domain"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
)

type AdminUseCase interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.AdminToken, error)
}
