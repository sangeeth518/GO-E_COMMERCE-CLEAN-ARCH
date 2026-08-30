package interfaces

import (
	"github.com/sangeeth518/go-Ecommerce/pkg/domain"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
)

type CartRepository interface {
	GetCartByUserID(userID int) (domain.Cart, error)
	CreateCart(userID int) (domain.Cart, error)
	GetCartItem(cartID, inventoryID int) (domain.CartItem, error)
	AddCartItem(item domain.CartItem) error
	UpdateCartItem(item domain.CartItem) error
	RemoveCartItem(cartID, inventoryID int) error
	GetCartItems(cartID int) ([]models.CartItemResponse, error)
	GetProductStockAndPrice(inventoryID int) (stock int, price float64, err error)
	UpdateCartTotal(cartID int, total float64) error
}
