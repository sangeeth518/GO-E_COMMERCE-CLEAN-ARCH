package domain

type Cart struct {
	Id         int        `json:"id" gorm:"primaryKey"`
	UserId     int        `json:"user_id" gorm:"not null;unique"`
	User       User       `json:"-" gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
	CartItems  []CartItem `json:"cart_items" gorm:"foreignKey:CartId"`
	TotalPrice float64    `json:"total_price" gorm:"default:0"`
}

type CartItem struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	CartId      int       `json:"cart_id" gorm:"not null;uniqueIndex:idx_cart_inventory"`
	Cart        Cart      `json:"-" gorm:"foreignKey:CartId;constraint:OnDelete:CASCADE"`
	InventoryId int       `json:"inventory_id" gorm:"not null;uniqueIndex:idx_cart_inventory"`
	Inventory   Inventory `json:"-" gorm:"foreignKey:InventoryId;constraint:OnDelete:CASCADE"`
	Quantity    int       `json:"quantity" gorm:"not null;default:1"`
	TotalPrice  float64   `json:"total_price" gorm:"not null;default:0"`
}
