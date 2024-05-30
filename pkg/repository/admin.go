package repository

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/sangeeth518/go-Ecommerce/pkg/domain"
	interfaces "github.com/sangeeth518/go-Ecommerce/pkg/repository/interface"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
	"gorm.io/gorm"
)

type Adminrepo struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepo {
	return &Adminrepo{
		DB: DB,
	}
}
func (ad *Adminrepo) LoginHandler(admindetails models.AdminLogin) (domain.Admin, error) {
	var admincomparedetails domain.Admin
	if err := ad.DB.Raw("select * from admins where email = ?", admindetails.Email).Scan(&admincomparedetails).Error; err != nil {
		return domain.Admin{}, err
	}
	return admincomparedetails, nil
}

func (ad *Adminrepo) GetUserById(id string) (domain.User, error) {
	user_id, err := strconv.Atoi(id)
	if err != nil {
		return domain.User{}, err
	}
	var count int
	if err := ad.DB.Raw("select count(*) from users where id =?", user_id).Scan(&count).Error; err != nil {
		return domain.User{}, err
	}
	if count < 1 {
		return domain.User{}, errors.New("user for the given id dosen't exist")
	}
	query := fmt.Sprintf("select * from users where id = %d", user_id)
	var userdetails domain.User
	if err := ad.DB.Raw(query).Scan(&userdetails).Error; err != nil {
		return domain.User{}, err
	}
	return userdetails, nil

}

//Func which will both block and unblock user

func (ad *Adminrepo) BlockUserById(user domain.User) error {
	err := ad.DB.Exec("update users set blocked =? where id =?", user.Blocked, user.Id).Error
	if err != nil {
		return err
	}
	return nil

}

func (ad *Adminrepo) GetUsers(page int, count int) ([]models.UserdetailsAtAdmin, error) {

	if page == 0 {
		page = 1
	}
	offset := (page - 1) * count

	var userdetails []models.UserdetailsAtAdmin

	if err := ad.DB.Raw("select id , email, phone,blocked from users limit ? offset ?", count, offset).Scan(&userdetails).Error; err != nil {
		return []models.UserdetailsAtAdmin{}, errors.New("err")
	}
	fmt.Printf("%d", offset)
	return userdetails, nil

}
