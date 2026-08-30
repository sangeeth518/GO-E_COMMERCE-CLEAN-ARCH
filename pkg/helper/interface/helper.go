package interfaces

import (
	"context"
	"mime/multipart"

	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
)

type Helper interface {
	GenerateTokenAdmin(admin models.AdminDetailResponse) (string, string, error)
	GenerateTokenClient(user models.UserDetailsResponse) (string, error)
	PasswordHashing(password string) (string, error)
	CompareHashPassword(password string, givenpass string) error
	AddProductImage(ctx context.Context, file *multipart.FileHeader, productId int) (string, error)
	GetPresignedURL(ctx context.Context, key string) (string, error)
	DeleteProductImageFromS3(ctx context.Context, key string) error
}


