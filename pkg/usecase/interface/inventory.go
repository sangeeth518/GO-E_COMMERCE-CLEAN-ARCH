package interfaces

import (
	"context"
	"mime/multipart"

	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
)

type InventoryUsecase interface {
	AddProduct(product models.AddProduct) (models.ProductResponse, error)
	ListProducts(page, limit int) ([]models.Inventories, error)
	AddProductImages(ctx context.Context, files []*multipart.FileHeader, productId int) ([]models.ImageUploadResult, error)
	GetProductByID(ctx context.Context, productId int) (models.ProductWithImages, error)
	DeleteProductImage(ctx context.Context, imageId int) error
	SetPrimaryImage(ctx context.Context, imageId int) error
	DeleteProduct(ctx context.Context, productId int) error
}
