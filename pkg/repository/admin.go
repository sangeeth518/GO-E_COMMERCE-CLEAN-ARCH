package repository

import (
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
	if err := ad.DB.Raw("select * from users where email = ?", admindetails.Email).Scan(&admincomparedetails).Error; err != nil {
		return domain.Admin{}, err
	}
	return admincomparedetails, nil
}
