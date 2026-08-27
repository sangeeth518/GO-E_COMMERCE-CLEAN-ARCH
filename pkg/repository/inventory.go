package repository

import (
	"gorm.io/gorm"

	interfaces "github.com/sangeeth518/go-Ecommerce/pkg/repository/interface"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
)

type inventoryRepo struct {
	DB *gorm.DB
}

func NewInventoryRepository(DB *gorm.DB) interfaces.InventoryRepo {
	return &inventoryRepo{
		DB: DB,
	}
}

func (i *inventoryRepo) AddProduct(product models.AddProduct) (models.ProductResponse, error) {
	var productResponse models.ProductResponse
	err := i.DB.Raw("insert into inventories (category_id,product_name,description,size,stock,price) values (?,?,?,?,?,?) returning id,product_name AS name , description,size,stock,price", product.CategoryId, product.ProductName, product.Description, product.Size, product.Stock, product.Price).Scan(&productResponse).Error
	if err != nil {

		return models.ProductResponse{}, err
	}
	return productResponse, nil
}

func (i *inventoryRepo) ListProducts(page, limit int) ([]models.Inventories, error) {
	var inventories []models.Inventories
	offset := (page - 1) * limit
	query := `
		SELECT 
			inventories.id,
			inventories.category_id,
			categories.name AS category,
			inventories.product_name,
			inventories.description,
			inventories.size,
			inventories.stock,
			inventories.price
		FROM inventories
		LEFT JOIN categories ON inventories.category_id = categories.id
		ORDER BY inventories.id ASC
		LIMIT ? OFFSET ?
	`
	if err := i.DB.Raw(query, limit, offset).Scan(&inventories).Error; err != nil {
		return nil, err
	}
	return inventories, nil
}

