package usecase

import (
	"errors"

	"github.com/jinzhu/copier"
	"github.com/sangeeth518/go-Ecommerce/pkg/domain"
	helper_interface "github.com/sangeeth518/go-Ecommerce/pkg/helper/interface"
	interfaces "github.com/sangeeth518/go-Ecommerce/pkg/repository/interface"
	services "github.com/sangeeth518/go-Ecommerce/pkg/usecase/interface"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
	"golang.org/x/crypto/bcrypt"
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

	//Getting Admin details based on the email provided
	admindetails, err := ad.adminrepository.LoginHandler(adminDetails)
	if err != nil {
		return domain.AdminToken{}, err
	}
	//compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(admindetails.Password), []byte(adminDetails.Password))
	if err != nil {
		return domain.AdminToken{}, err
	}
	var adminDetailResponse models.AdminDetailResponse
	err = copier.Copy(&adminDetailResponse, &admindetails)
	if err != nil {
		return domain.AdminToken{}, err
	}
	access, refresh, err := ad.helper.GenerateTokenAdmin(adminDetailResponse)

	if err != nil {
		return domain.AdminToken{}, err
	}

	return domain.AdminToken{
		Admin:   adminDetailResponse,
		Token:   access,
		Refresh: refresh,
	}, nil

}

func (au *adminUsecase) BlockUser(id string) error {
	userdetails, err := au.adminrepository.GetUserById(id)
	if err != nil {
		return err
	}
	if userdetails.Blocked {
		return errors.New("already blocked")
	} else {
		userdetails.Blocked = true
	}

	err = au.adminrepository.BlockUserById(userdetails)
	if err != nil {
		return err
	}
	return nil
}

func (au *adminUsecase) UnblockUser(id string) error {
	user, err := au.adminrepository.GetUserById(id)
	if err != nil {
		return err
	}

	if user.Blocked {
		user.Blocked = false
	} else {
		return errors.New("user already unblocked")
	}

	err = au.adminrepository.BlockUserById(user)
	if err != nil {
		return err
	}
	return nil
}
