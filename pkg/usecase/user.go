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
		return models.UserToken{}, errors.New("user already exist")
	}
	//comparing password & confirmpassword
	if user.Password != user.ConfirmPassword {
		return models.UserToken{}, errors.New("password dosen't match")
	}
	//passwordHashing using bcrypt
	hashedpassword, err := use.helper.PasswordHashing(user.Password)
	if err != nil {
		return models.UserToken{}, errors.New("password hashing error")
	}
	user.Password = hashedpassword

	userdata, err := use.userrepo.UserSignup(user)
	if err != nil {
		return models.UserToken{}, errors.New("couldn't add user")
	}
	token, err := use.helper.GenerateTokenClient(userdata)
	if err != nil {
		return models.UserToken{}, errors.New("couldn't create token due to internal error")
	}

	return models.UserToken{
		User:  userdata,
		Token: token,
	}, nil

}

func (use *userUsecase) UserLogin(user models.UserLogin) (models.UserToken, error) {
	//Check the user exist or not
	ok := use.userrepo.CheckUserAvailability(user.Email)
	if !ok {
		return models.UserToken{}, errors.New("user didn't exist")
	}
	//check admin blocked this user or not
	permission, err := use.userrepo.UserBlockStatus(user.Email)
	if err != nil {
		return models.UserToken{}, errors.New("internal server error")
	}
	if permission {
		return models.UserToken{}, errors.New("usere is blocked by admin")
	}

	//Get the user details from the given email
	user_details, err := use.userrepo.FindUserByEmail(user)
	if err != nil {
		return models.UserToken{}, errors.New("internal server error")
	}

	// comapre passwords
	err = use.helper.CompareHashPassword(user_details.Password, user.Password)
	if err != nil {
		return models.UserToken{}, errors.New("password Incorrect")
	}
	var userresponse models.UserDetailsResponse
	userresponse.Id = int(user_details.Id)
	userresponse.Name = user_details.Name
	userresponse.Email = user_details.Email
	userresponse.Phone = user_details.Phone

	//Generate token
	tokenstring, err := use.helper.GenerateTokenClient(userresponse)
	if err != nil {
		return models.UserToken{}, errors.New("could'nt create token for users")
	}
	return models.UserToken{
		User:  userresponse,
		Token: tokenstring,
	}, nil
}

func (u *userUsecase) AddAdress(id int, adress models.AddAdress) error {
	err := u.userrepo.AddAdress(id, adress)
	if err != nil {
		return err
	}
	return nil
}

func (uc *userUsecase) ChangePassword(id int, password string, newpass string, confrmpass string) error {
	user_password, err := uc.userrepo.GetPassword(id)
	if err != nil {
		return errors.New("internal error")
	}

	err = uc.helper.CompareHashPassword(user_password, password)
	if err != nil {
		return errors.New("password incorrect")
	}
	if newpass != confrmpass {
		return errors.New("passwords does not match")
	}

	new_pass, err := uc.helper.PasswordHashing(newpass)
	if err != nil {
		return errors.New("error in password hashing")
	}
	return uc.userrepo.Changepassword(id, string(new_pass))

}
