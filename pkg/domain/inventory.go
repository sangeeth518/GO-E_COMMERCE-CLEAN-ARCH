package domain

type Category struct {
	Id   int    `json:"id" gorm:"unique;not null"`
	Name string `json:"name"`
}

type Inventories struct {
	Id          int      `json:"id" gorm:"unique;not null" `
	CategoryId  int      `json:"categroy_id"`
	Category    Category `json:"-" foreignkey:"CategoryId"`
	ProductName string   `json:"product_name"`
	Image       string   `json:"image"`
	Size        string   `json:"size" gorm:"size:5;default:'M';check:size IN ('S', 'M', 'L', 'XL')"`
	Stock       int      `json:"stock"`
	Price       float64  `json:"price"`
}
