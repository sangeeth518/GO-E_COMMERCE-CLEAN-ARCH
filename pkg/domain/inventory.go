package domain

type Category struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type Inventory struct {
	Id          int            `json:"id" gorm:"primaryKey" `
	CategoryId  int            `json:"category_id"`
	Category    Category       `json:"-" gorm:"foreignKey:CategoryId"`
	ProductName string         `json:"product_name"`
	Description string         `json:"description"`
	Size        string         `json:"size" gorm:"size:5;default:'M';check:size IN ('S', 'M', 'L', 'XL')"`
	Stock       int            `json:"stock"`
	Price       float64        `json:"price"`
	Images      []ProductImage `json:"images" gorm:"foreignKey:ProductID"`
}

type ProductImage struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	ProductID int    `json:"product_id"`
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary" gorm:"default:false"`
}
