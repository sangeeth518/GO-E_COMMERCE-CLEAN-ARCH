package interfaces

import "github.com/sangeeth518/go-Ecommerce/pkg/domain"

type CateoryUsecase interface {
	AddCategory(category domain.Category) (domain.Category, error)
}
