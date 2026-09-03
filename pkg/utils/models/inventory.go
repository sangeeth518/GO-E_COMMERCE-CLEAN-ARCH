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

// response for single product with images
type ProductWithImages struct {
	Id          int                    `json:"id"`
	CategoryId  int                    `json:"category_id"`
	Category    string                 `json:"category"`
	ProductName string                 `json:"product_name"`
	Description string                 `json:"description"`
	Size        string                 `json:"size"`
	Stock       int                    `json:"stock"`
	Price       float64                `json:"price"`
	Images      []ProductImageResponse `json:"images"`
}

type ProductImageResponse struct {
	Id        int    `json:"id"`
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

type ImageUploadResult struct {
	Index        int    `json:"index"`
	Filename     string `json:"filename"`
	Success      bool   `json:"success"`
	PresignedURL string `json:"presigned_url,omitempty"`
	Error        string `json:"error,omitempty"`
}
