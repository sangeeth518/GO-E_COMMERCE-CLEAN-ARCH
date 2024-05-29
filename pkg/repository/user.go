package repository

import (
	"errors"

	interfaces "github.com/sangeeth518/go-Ecommerce/pkg/repository/interface"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepo {
	return &userRepository{
		DB: DB,
	}
}

func (ur *userRepository) CheckUserAvailability(email string) bool {
	var count int
	if err := ur.DB.Raw("select count(*) from users where email =?", email).Scan(&count).Error; err != nil {
		return false
	}
	//if count is greater than 0 that means user with same email id already exist and it returns true
	return count > 0

}

func (ur *userRepository) UserBlockStatus(emil string) (bool, error) {
	var permission bool
	err := ur.DB.Raw("SELECT blocked FROM users WHERE email = $1", emil).Scan(&permission).Error
	if err != nil {
		return false, err
	}
	return permission, nil
}

func (ur *userRepository) FindUserByEmail(user models.UserLogin) (models.UserSigninResponse, error) {
	var userResponse models.UserSigninResponse
	if err := ur.DB.Raw("SELECT * FROM users WHERE email =$1 and blocked = false", user.Email).Scan(&userResponse).Error; err != nil {
		return models.UserSigninResponse{}, errors.New("error checking user details")
	}
	return userResponse, nil

}

func (ur *userRepository) UserSignup(user models.UserDetails) (models.UserDetailsResponse, error) {
	var userdetails models.UserDetailsResponse
	if err := ur.DB.Raw("INSERT INTO users (name, email, password, phone) values (?, ?, ?, ?)  RETURNING id, name, email, phone", user.Name, user.Email, user.Password, user.Phone).Scan(&userdetails).Error; err != nil {
		return models.UserDetailsResponse{}, err
	}
	return userdetails, nil

}

func (ur *userRepository) Changepassword(id int, passowrd string) error {
	err := ur.DB.Exec("update users set password = ? where id = ?", id, passowrd).Error
	if err != nil {

		return err
	}
	return nil

}

func (u *userRepository) GetPassword(id int) (string, error) {
	var userpassword string
	err := u.DB.Raw("select password from users where id = ?", id).Scan(&userpassword).Error
	if err != nil {
		return "", err
	}
	return userpassword, nil
}
