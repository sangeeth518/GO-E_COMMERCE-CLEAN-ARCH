package repository

import (
	"errors"

	"github.com/sangeeth518/go-Ecommerce/pkg/domain"
	interfaces "github.com/sangeeth518/go-Ecommerce/pkg/repository/interface"
	"gorm.io/gorm"
)

type categoryRepo struct {
	DB *gorm.DB
}

func NewCategoryRepository(DB *gorm.DB) interfaces.CategoryRepo {
	return &categoryRepo{
		DB: DB,
	}
}

func (c *categoryRepo) AddCategory(category domain.Category) (domain.Category, error) {
	var cat domain.Category

	if category.Name == "" {
		return domain.Category{}, errors.New("category name cannot be empty")
	}
	if err := c.DB.Raw("insert into categories (name) values (?) returning name , id ", category.Name).Scan(&cat).Error; err != nil {
		return domain.Category{}, err
	}
	return cat, nil
}
