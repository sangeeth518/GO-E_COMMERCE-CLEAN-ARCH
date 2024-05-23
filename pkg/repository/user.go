package repository

import (
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

}

func (ur *userRepository) UserSignup(user models.UserDetails) (models.UserDetailsResponse, error) {
	var userdetails models.UserDetailsResponse
	if err := ur.DB.Raw("INSERT INTO users (name, email, password, phone) values (?, ?, ?, ?)  RETURNING id, name, email, phone", user.Name, user.Email, user.Password, user.Phone).Scan(&userdetails).Error; err != nil {
		return models.UserDetailsResponse{}, err
	}
	return userdetails, nil

}
