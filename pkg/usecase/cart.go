package usecase

import (
	"errors"

	"github.com/sangeeth518/go-Ecommerce/pkg/domain"
	repoInterfaces "github.com/sangeeth518/go-Ecommerce/pkg/repository/interface"
	interfaces "github.com/sangeeth518/go-Ecommerce/pkg/usecase/interface"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
	"gorm.io/gorm"
)

type CartUsecase struct {
	cartRepo repoInterfaces.CartRepository
}

func NewCartUsecase(cartRepo repoInterfaces.CartRepository) interfaces.CartUsecase {
	return &CartUsecase{cartRepo: cartRepo}
}

func (u *CartUsecase) refreshCartTotal(cartID int) (float64, error) {
	items, err := u.cartRepo.GetCartItems(cartID)
	if err != nil {
		return 0, err
	}
	var grandTotal float64
	for _, item := range items {
		grandTotal += item.TotalPrice
	}
	err = u.cartRepo.UpdateCartTotal(cartID, grandTotal)
	return grandTotal, err
}

func (u *CartUsecase) AddToCart(userID int, req models.AddToCart) error {
	stock, price, err := u.cartRepo.GetProductStockAndPrice(req.InventoryId)
	if err != nil {
		return errors.New("product does not exist")
	}

	if req.Quantity > stock {
		return errors.New("insufficient product stock available")
	}

	cart, err := u.cartRepo.GetCartByUserID(userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		cart, err = u.cartRepo.CreateCart(userID)
		if err != nil {
			return errors.New("could not create cart")
		}
	} else if err != nil {
		return err
	}

	existingItem, err := u.cartRepo.GetCartItem(cart.Id, req.InventoryId)
	if err == nil {
		newQuantity := existingItem.Quantity + req.Quantity
		if newQuantity > stock {
			return errors.New("insufficient product stock available")
		}
		existingItem.Quantity = newQuantity
		existingItem.TotalPrice = float64(newQuantity) * price
		if err := u.cartRepo.UpdateCartItem(existingItem); err != nil {
			return err
		}
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		newItem := domain.CartItem{
			CartId:      cart.Id,
			InventoryId: req.InventoryId,
			Quantity:    req.Quantity,
			TotalPrice:  float64(req.Quantity) * price,
		}
		if err := u.cartRepo.AddCartItem(newItem); err != nil {
			return err
		}
	} else {
		return err
	}

	_, err = u.refreshCartTotal(cart.Id)
	return err
}

func (u *CartUsecase) ViewCart(userID int) (models.CartResponse, error) {
	cart, err := u.cartRepo.GetCartByUserID(userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.CartResponse{UserId: userID, Items: []models.CartItemResponse{}, GrandTotal: 0}, nil
	} else if err != nil {
		return models.CartResponse{}, err
	}

	items, err := u.cartRepo.GetCartItems(cart.Id)
	if err != nil {
		return models.CartResponse{}, err
	}

	var grandTotal float64
	for _, item := range items {
		grandTotal += item.TotalPrice
	}

	return models.CartResponse{
		CartId:     cart.Id,
		UserId:     userID,
		Items:      items,
		GrandTotal: grandTotal,
	}, nil
}

func (u *CartUsecase) UpdateQuantity(userID int, req models.UpdateQuantityReq) error {
	cart, err := u.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return errors.New("cart not found")
	}

	stock, price, err := u.cartRepo.GetProductStockAndPrice(req.InventoryId)
	if err != nil {
		return errors.New("product does not exist")
	}

	if req.Quantity > stock {
		return errors.New("insufficient stock")
	}

	cartItem, err := u.cartRepo.GetCartItem(cart.Id, req.InventoryId)
	if err != nil {
		return errors.New("item not present in cart")
	}

	cartItem.Quantity = req.Quantity
	cartItem.TotalPrice = float64(req.Quantity) * price

	if err := u.cartRepo.UpdateCartItem(cartItem); err != nil {
		return err
	}

	_, err = u.refreshCartTotal(cart.Id)
	return err
}

func (u *CartUsecase) RemoveProductFromCart(userID, inventoryID int) error {
	cart, err := u.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return errors.New("cart not found")
	}

	if err := u.cartRepo.RemoveCartItem(cart.Id, inventoryID); err != nil {
		return err
	}

	_, err = u.refreshCartTotal(cart.Id)
	return err
}
