package domain

import "time"

type Order struct {
	Id                int         `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId            int         `json:"user_id" gorm:"not null"`
	User              User        `json:"-" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
	AddressId         int         `json:"address_id" gorm:"not null"`
	Address           Adress      `json:"address" gorm:"foreignKey:AddressId"`
	PaymentMethod     string      `json:"payment_method" gorm:"not null"`
	PaymentStatus     string      `json:"payment_status" gorm:"default:'pending'"`
	OrderStatus       string      `json:"order_status" gorm:"default:'pending'"`
	FinalPrice        float64     `json:"final_price" gorm:"not null"`
	Discount          float64     `json:"discount" gorm:"default:0"`
	RazorpayOrderId   string      `json:"razorpay_order_id" gorm:"default:''"`
	RazorpayPaymentId string      `json:"razorpay_payment_id" gorm:"default:''"`
	CreatedAt         time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt         time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
	OrderItems        []OrderItem `json:"order_items" gorm:"foreignKey:OrderId;constraint:OnDelete:CASCADE"`
}

type OrderItem struct {
	Id          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderId     int       `json:"order_id" gorm:"not null"`
	Order       Order     `json:"-" gorm:"foreignKey:OrderId;constraint:OnDelete:CASCADE"`
	InventoryId int       `json:"inventory_id" gorm:"not null"`
	Inventory   Inventory `json:"inventory" gorm:"foreignKey:InventoryId"`
	Quantity    int       `json:"quantity" gorm:"not null;default:1"`
	UnitPrice   float64   `json:"unit_price" gorm:"not null"`
	TotalPrice  float64   `json:"total_price" gorm:"not null"`
}
