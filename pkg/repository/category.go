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

func (c *categoryRepo) ShowCategories() ([]domain.Category, error) {
	var cat []domain.Category
	if err := c.DB.Raw("select * from categories").Scan(&cat).Error; err != nil {
		return []domain.Category{}, err
	}
	return cat, nil

}

func (c *categoryRepo) EditCategory(category domain.Category) error {
	if err := c.DB.Model(&category).Where("id=?", category.Id).Update("name", category.Name).Error; err != nil {
		return err
	}
	return nil

}
