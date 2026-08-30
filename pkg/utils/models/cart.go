package models

type AddToCart struct {
	InventoryId int `json:"inventory_id" validate:"required"`
	Quantity    int `json:"quantity" validate:"required,min=1"`
}

type UpdateQuantityReq struct {
	InventoryId int `json:"inventory_id" validate:"required"`
	Quantity    int `json:"quantity" validate:"required,min=1"`
}

type CartItemResponse struct {
	InventoryId int     `json:"inventory_id"`
	ProductName string  `json:"product_name"`
	Size        string  `json:"size"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	TotalPrice  float64 `json:"total_price"`
}

type CartResponse struct {
	CartId     int                `json:"cart_id"`
	UserId     int                `json:"user_id"`
	Items      []CartItemResponse `json:"items"`
	GrandTotal float64            `json:"grand_total"`
}
