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
		SELECT 	inventories.id,
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

func (i *inventoryRepo) AddProductImage(productId int, imageUrl string, isPrimary bool) error {
	err := i.DB.Exec("insert into product_images (product_id,image_url,is_primary) values(?,?,?)", productId, imageUrl, isPrimary).Error
	if err != nil {
		return err
	}
	return nil

}

func (i *inventoryRepo) HasPrimaryImage(productId int) (bool, error) {
	var count int
	if err := i.DB.Raw("select count(*) from product_images where product_id =? and is_primary=true", productId).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (i *inventoryRepo) CheckProductExists(productId int) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM inventories WHERE id = ?`
	if err := i.DB.Raw(query, productId).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (i *inventoryRepo) GetProductByID(productId int) (models.Inventories, error) {
	var product models.Inventories
	query := `
		SELECT 	inventories.id,
			inventories.category_id,
			categories.name AS category,
			inventories.product_name,
			inventories.description,
			inventories.size,
			inventories.stock,
			inventories.price
		FROM inventories
		LEFT JOIN categories ON inventories.category_id = categories.id
		WHERE inventories.id = ?
	`
	if err := i.DB.Raw(query, productId).Scan(&product).Error; err != nil {
		return models.Inventories{}, err
	}
	return product, nil
}

func (i *inventoryRepo) GetProductImages(productId int) ([]models.ProductImageResponse, error) {
	var images []models.ProductImageResponse
	query := `SELECT id, image_url, is_primary FROM product_images WHERE product_id = ? ORDER BY is_primary DESC, id ASC`
	if err := i.DB.Raw(query, productId).Scan(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}

func (i *inventoryRepo) GetImageURLByID(imageId int) (string, error) {
	var imageURL string
	query := `SELECT image_url FROM product_images WHERE id = ?`
	if err := i.DB.Raw(query, imageId).Scan(&imageURL).Error; err != nil {
		return "", err
	}
	return imageURL, nil
}

func (i *inventoryRepo) DeleteProductImageByID(imageId int) error {
	query := `DELETE FROM product_images WHERE id = ?`
	return i.DB.Exec(query, imageId).Error
}

func (i *inventoryRepo) GetImageURLsByProductID(productId int) ([]string, error) {
	var urls []string
	query := `SELECT image_url FROM product_images WHERE product_id = ?`
	if err := i.DB.Raw(query, productId).Scan(&urls).Error; err != nil {
		return nil, err
	}
	return urls, nil
}

func (i *inventoryRepo) DeleteAllProductImages(productId int) error {
	query := `DELETE FROM product_images WHERE product_id = ?`
	return i.DB.Exec(query, productId).Error
}

func (i *inventoryRepo) DeleteProduct(productId int) error {
	query := `DELETE FROM inventories WHERE id = ?`
	return i.DB.Exec(query, productId).Error
}

