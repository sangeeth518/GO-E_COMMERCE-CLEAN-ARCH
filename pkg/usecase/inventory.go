package usecase

import (
	"context"
	"errors"
	"mime/multipart"
	"sync"

	helper_interface "github.com/sangeeth518/go-Ecommerce/pkg/helper/interface"
	interfaces "github.com/sangeeth518/go-Ecommerce/pkg/repository/interface"
	services "github.com/sangeeth518/go-Ecommerce/pkg/usecase/interface"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
)

type inventoryUsecase struct {
	invrepo interfaces.InventoryRepo
	helper  helper_interface.Helper
}

func NewInventoryUsecase(invrepo interfaces.InventoryRepo, helper helper_interface.Helper) services.InventoryUsecase {
	return &inventoryUsecase{
		invrepo: invrepo,
		helper:  helper}
}

func (i *inventoryUsecase) AddProduct(product models.AddProduct) (models.ProductResponse, error) {
	if product.Price <= 0 {
		return models.ProductResponse{}, errors.New("price must be greater than zero")
	}
	if product.Stock < 0 {
		return models.ProductResponse{}, errors.New("stock cannot be negative")
	}
	productresponse, err := i.invrepo.AddProduct(product)
	if err != nil {
		return models.ProductResponse{}, errors.New("failed to add in Db")
	}
	return productresponse, nil

}

func (i *inventoryUsecase) ListProducts(page, limit int) ([]models.Inventories, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	productDetails, err := i.invrepo.ListProducts(page, limit)
	if err != nil {
		return nil, err
	}
	return productDetails, nil
}

//Add Product Images

func (i *inventoryUsecase) AddProductImages(ctx context.Context, files []*multipart.FileHeader, productId int) ([]models.ImageUploadResult, error) {
	exists, err := i.invrepo.CheckProductExists(productId)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.New("product does not exist")
	}

	hasPrimary, err := i.invrepo.HasPrimaryImage(productId)
	if err != nil {
		return nil, errors.New("cannot check the existing images")
	}

	results := make([]models.ImageUploadResult, len(files))
	var wg sync.WaitGroup

	for idx, f := range files {
		index, file := idx, f
		wg.Add(1)

		go func() {
			defer wg.Done()
			results[index] = models.ImageUploadResult{
				Index:    index,
				Filename: file.Filename,
			}
			// 1. Upload to S3 and get the S3 key
			key, err := i.helper.AddProductImage(ctx, file, productId)
			if err != nil {
				results[index].Error = "upload failed " + err.Error()
				return
			}

			isprimary := false
			if !hasPrimary && index == 0 {
				isprimary = true
			}

			// 2. Save S3 key in database
			err = i.invrepo.AddProductImage(productId, key, isprimary)
			if err != nil {
				results[index].Error = "failed to save image key in database" + err.Error()
				return
			}

			// 3. Generate presigned URL for response
			presignedURL, err := i.helper.GetPresignedURL(ctx, key)
			if err != nil {
				results[index].Error = "Upload success but failed to generate preview URL"
				return
			}
			results[index].PresignedURL = presignedURL
			results[index].Success = true
		}()
	}

	wg.Wait()
	return results, nil
}

func (i *inventoryUsecase) GetProductByID(ctx context.Context, productId int) (models.ProductWithImages, error) {
	exists, err := i.invrepo.CheckProductExists(productId)
	if err != nil {
		return models.ProductWithImages{}, err
	}
	if !exists {
		return models.ProductWithImages{}, errors.New("product does not exist")
	}

	product, err := i.invrepo.GetProductByID(productId)
	if err != nil {
		return models.ProductWithImages{}, err
	}

	images, err := i.invrepo.GetProductImages(productId)
	if err != nil {
		return models.ProductWithImages{}, err
	}

	// Generate presigned URL for each image key stored in the database
	for idx := range images {
		presignedURL, err := i.helper.GetPresignedURL(ctx, images[idx].ImageUrl)
		if err == nil {
			images[idx].ImageUrl = presignedURL
		}
	}

	return models.ProductWithImages{
		Id:          int(product.ID),
		CategoryId:  product.CategoryID,
		Category:    product.Category,
		ProductName: product.ProductName,
		Description: product.Description,
		Size:        product.Size,
		Stock:       product.Stock,
		Price:       product.Price,
		Images:      images,
	}, nil
}

func (i *inventoryUsecase) DeleteProductImage(ctx context.Context, imageId int) error {
	image, err := i.invrepo.GetImageByID(imageId)
	if err != nil {
		return errors.New("image not found in database")
	}

	// Delete from S3 using key directly
	if err := i.helper.DeleteProductImageFromS3(ctx, image.ImageUrl); err != nil {
		return errors.New("failed to delete image from S3: " + err.Error())
	}

	// Delete from DB
	if err := i.invrepo.DeleteProductImageByID(imageId); err != nil {
		return errors.New("failed to delete image from database")
	}

	// If the deleted image was the primary image, promote the next remaining image to primary
	if image.IsPrimary {
		_ = i.invrepo.SetFirstImageAsPrimary(image.ProductID)
	}

	return nil
}

func (i *inventoryUsecase) SetPrimaryImage(ctx context.Context, imageId int) error {
	image, err := i.invrepo.GetImageByID(imageId)
	if err != nil {
		return errors.New("image not found in database")
	}

	if err := i.invrepo.SetImageAsPrimary(image.ProductID, imageId); err != nil {
		return errors.New("failed to set image as primary")
	}

	return nil
}

func (i *inventoryUsecase) DeleteProduct(ctx context.Context, productId int) error {
	exists, err := i.invrepo.CheckProductExists(productId)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("product does not exist")
	}

	keys, err := i.invrepo.GetImageURLsByProductID(productId)
	if err != nil {
		return errors.New("failed to fetch product images")
	}

	// Delete each object from S3 using key directly
	for _, key := range keys {
		_ = i.helper.DeleteProductImageFromS3(ctx, key)
	}

	if err := i.invrepo.DeleteAllProductImages(productId); err != nil {
		return errors.New("failed to delete product images from database")
	}

	if err := i.invrepo.DeleteProduct(productId); err != nil {
		return errors.New("failed to delete product from database")
	}

	return nil
}
