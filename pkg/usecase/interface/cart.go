package interfaces

import "github.com/sangeeth518/go-Ecommerce/pkg/utils/models"

type CartUsecase interface {
	AddToCart(userID int, req models.AddToCart) error
	ViewCart(userID int) (models.CartResponse, error)
	UpdateQuantity(userID int, req models.UpdateQuantityReq) error
	RemoveProductFromCart(userID, inventoryID int) error
}
