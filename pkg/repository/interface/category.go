package interfaces

import "github.com/sangeeth518/go-Ecommerce/pkg/domain"

type CategoryRepo interface {
	AddCategory(category domain.Category) (domain.Category, error)
}
