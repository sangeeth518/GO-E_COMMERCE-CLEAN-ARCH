package models

type AddProduct struct {
	CategoryId  int     `json:"-"`
	ProductName string  `json:"product_name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Size        string  `json:"size" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

type ProductResponse struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Size        string  `json:"size"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}

type Inventories struct {
	ID          uint    `json:"id"`
	CategoryID  int     `json:"category_id"`
	Category    string  `json:"category"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Size        string  `json:"size"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}

