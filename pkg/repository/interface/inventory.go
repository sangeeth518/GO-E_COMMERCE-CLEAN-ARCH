package interfaces

import "github.com/sangeeth518/go-Ecommerce/pkg/utils/models"

type InventoryRepo interface {
	AddProduct(product models.AddProduct) (models.ProductResponse, error)
	ListProducts(page, limit int) ([]models.Inventories, error)
	AddProductImage(productId int, imageUrl string, isPrimary bool) error
	HasPrimaryImage(productId int) (bool, error)
	CheckProductExists(productId int) (bool, error)
	GetProductByID(productId int) (models.Inventories, error)
	GetProductImages(productId int) ([]models.ProductImageResponse, error)
	GetImageURLByID(imageId int) (string, error)
	DeleteProductImageByID(imageId int) error
	GetImageURLsByProductID(productId int) ([]string, error)
	DeleteAllProductImages(productId int) error
	DeleteProduct(productId int) error
}

