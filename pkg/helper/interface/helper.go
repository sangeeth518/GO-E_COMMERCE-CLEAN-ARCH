package interfaces

import "github.com/sangeeth518/go-Ecommerce/pkg/utils/models"

type Helper interface {
	GenerateTokenAdmin(admin models.AdminDetailResponse) (string, string, error)
}
