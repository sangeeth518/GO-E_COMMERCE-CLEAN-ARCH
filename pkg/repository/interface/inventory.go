package interfaces

import "github.com/sangeeth518/go-Ecommerce/pkg/utils/models"

type InventoryRepo interface {
	AddProduct(product models.AddProduct) (models.ProductResponse, error)
	ListProducts(page, limit int) ([]models.Inventories, error)
	AddProductImage(productId int, imageUrl string, isPrimary bool) error
	HasPrimaryImage(productId int) (bool, error)
}
