package interfaces

import "github.com/sangeeth518/go-Ecommerce/pkg/utils/models"

type InventoryUsecase interface {
	AddProduct(product models.AddProduct) (models.ProductResponse, error)
	ListProducts(page, limit int) ([]models.Inventories, error)
}

