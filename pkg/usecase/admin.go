package usecase

import (
	"github.com/sangeeth518/go-Ecommerce/pkg/domain"
	helper_interface "github.com/sangeeth518/go-Ecommerce/pkg/helper/interface"
	interfaces "github.com/sangeeth518/go-Ecommerce/pkg/repository/interface"
	services "github.com/sangeeth518/go-Ecommerce/pkg/usecase/interface"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
)

type adminUsecase struct {
	adminrepository interfaces.AdminRepo
	helper          helper_interface.Helper
}

func NewAdminUsecase(repo interfaces.AdminRepo, h helper_interface.Helper) services.AdminUseCase {
	return &adminUsecase{
		adminrepository: repo,
		helper:          h,
	}
}

func (ad *adminUsecase) LoginHandler(adminDetails models.AdminLogin) (domain.AdminToken, error) {

}
