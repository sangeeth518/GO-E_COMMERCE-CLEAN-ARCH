package usecase

import (
	"github.com/sangeeth518/go-Ecommerce/pkg/domain"
	helper_interface "github.com/sangeeth518/go-Ecommerce/pkg/helper/interface"
	interfaces "github.com/sangeeth518/go-Ecommerce/pkg/repository/interface"
	services "github.com/sangeeth518/go-Ecommerce/pkg/usecase/interface"
)

type categoryUsecae struct {
	categoryrepo interfaces.CategoryRepo
	helper       helper_interface.Helper
}

func NewCategoryUsecase(category interfaces.CategoryRepo, helper helper_interface.Helper) services.CateoryUsecase {
	return &categoryUsecae{
		categoryrepo: category,
		helper:       helper,
	}
}

func (cu *categoryUsecae) AddCategory(category domain.Category) (domain.Category, error) {
	categoryresponse, err := cu.categoryrepo.AddCategory(category)
	if err != nil {
		return domain.Category{}, err
	}
	return categoryresponse, nil
}
