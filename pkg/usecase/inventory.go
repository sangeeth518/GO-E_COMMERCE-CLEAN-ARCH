package usecase

import (
	"context"
	"errors"
	"mime/multipart"

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
	if product.Price <= 0 {
		return models.ProductResponse{}, errors.New("price must be greater than zero")
	}
	if product.Stock < 0 {
		return models.ProductResponse{}, errors.New("stock cannot be negative")
	}
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

//Add Product Images

func (i *inventoryUsecase) AddProductImages(ctx context.Context, files []*multipart.FileHeader, productId int) ([]string, error) {

	hasPrimary, err := i.invrepo.HasPrimaryImage(productId)
	if err != nil {
		return nil, errors.New("cannot check the existing images")
	}

	var uploadedURLS []string

	for index, file := range files {
		url, err := i.helper.AddProductImage(ctx, file, productId)
		if err != nil {
			return nil, errors.New("failed to upload image: " + err.Error())
		}

		isprimary := false
		if !hasPrimary && index == 0 {
			isprimary = true
		}
		err = i.invrepo.AddProductImage(productId, url, isprimary)
		if err != nil {
			return nil, errors.New("Failed to upload image")
		}
		uploadedURLS = append(uploadedURLS, url)
	}
	return uploadedURLS, nil
}
