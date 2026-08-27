package usecase

import (
	"errors"

	helper_interface "github.com/sangeeth518/go-Ecommerce/pkg/helper/interface"
	interfaces "github.com/sangeeth518/go-Ecommerce/pkg/repository/interface"
	services "github.com/sangeeth518/go-Ecommerce/pkg/usecase/interface"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
)

type inventoryUsecase struct {
	invrepo interfaces.InventoryRepo
	helper  helper_interface.Helper
}

func NewInventoryUsecase(invrepo interfaces.InventoryRepo, helper helper_interface.Helper) services.InventoryUsecase {
	return &inventoryUsecase{
		invrepo: invrepo,
		helper:  helper}
}

func (i *inventoryUsecase) AddProduct(product models.AddProduct) (models.ProductResponse, error) {
	productresponse, err := i.invrepo.AddProduct(product)
	if err != nil {
		return models.ProductResponse{}, errors.New("failed to add in Db")
	}
	return productresponse, nil

}

func (i *inventoryUsecase) ListProducts(page, limit int) ([]models.Inventories, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	productDetails, err := i.invrepo.ListProducts(page, limit)
	if err != nil {
		return nil, err
	}
	return productDetails, nil
}

