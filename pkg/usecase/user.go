package usecase

import (
	"errors"

	helper_interface "github.com/sangeeth518/go-Ecommerce/pkg/helper/interface"
	interfaces "github.com/sangeeth518/go-Ecommerce/pkg/repository/interface"
	service "github.com/sangeeth518/go-Ecommerce/pkg/usecase/interface"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
)

type userUsecase struct {
	userrepo interfaces.UserRepo
	helper   helper_interface.Helper
}

func NewUserUsecase(user interfaces.UserRepo, h helper_interface.Helper) service.UserUsecase {
	return &userUsecase{
		userrepo: user,
		helper:   h,
	}
}

func (use *userUsecase) UserSignup(user models.UserDetails) (models.UserToken, error) {
	//checking whether the user with same email id already exist
	userexist := use.userrepo.CheckUserAvailability(user.Email)
	if userexist {
		return models.UserToken{}, errors.New("User already exist")
	}
	//comparing password & confirmpassword
	if user.Password != user.ConfirmPassword {
		return models.UserToken{}, errors.New("password dosen't match")
	}
	//passwordHashing using bcrypt

}
