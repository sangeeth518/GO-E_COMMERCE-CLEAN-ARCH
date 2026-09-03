package models

import "time"

// OrderIncoming represents user checkout request payload
type OrderIncoming struct {
	AddressId     int    `json:"address_id" validate:"required"`
	PaymentMethod string `json:"payment_method" validate:"required"` // e.g. "COD", "RAZORPAY"
}

// OrderResponse represents summary of an order
type OrderResponse struct {
	Id              int       `json:"id"`
	UserId          int       `json:"user_id"`
	AddressId       int       `json:"address_id"`
	PaymentMethod   string    `json:"payment_method"`
	PaymentStatus   string    `json:"payment_status"`
	OrderStatus     string    `json:"order_status"`
	FinalPrice      float64   `json:"final_price"`
	Discount        float64   `json:"discount"`
	RazorpayOrderId string    `json:"razorpay_order_id,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
}

// OrderItemResponse represents individual product line-item in an order
type OrderItemResponse struct {
	InventoryId int     `json:"inventory_id"`
	ProductName string  `json:"product_name"`
	Size        string  `json:"size"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	TotalPrice  float64 `json:"total_price"`
}

// OrderDetailsResponse represents complete order details along with items & address
type OrderDetailsResponse struct {
	Order      OrderResponse       `json:"order"`
	Address    AddressResponse     `json:"address"`
	OrderItems []OrderItemResponse `json:"order_items"`
}

// AddressResponse represents shipping address details inside order response
type AddressResponse struct {
	Name      string `json:"name"`
	HouseName string `json:"house_name"`
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
	Phone     string `json:"phone"`
	Pin       string `json:"pin"`
}

// summary of order — used in list view
type MyOrdersResponse struct {
	Id            int       `json:"id"`
	FinalPrice    float64   `json:"final_price"`
	PaymentMethod string    `json:"payment_method"`
	PaymentStatus string    `json:"payment_status"`
	OrderStatus   string    `json:"order_status"`
	CreatedAt     time.Time `json:"created_at"`
	ItemCount     int       `json:"item_count"`
}
