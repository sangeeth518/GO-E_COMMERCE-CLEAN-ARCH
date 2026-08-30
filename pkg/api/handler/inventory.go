package handler

import (
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	interfaces "github.com/sangeeth518/go-Ecommerce/pkg/usecase/interface"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/response"
)

type InventoryHandler struct {
	inv interfaces.InventoryUsecase
}

func NewInventoryHandler(inv interfaces.InventoryUsecase) *InventoryHandler {
	return &InventoryHandler{
		inv: inv,
	}
}

func (i *InventoryHandler) AddProduct(c *gin.Context) {

	catIDstr := c.Param("category_id")
	catId, err := strconv.Atoi(catIDstr)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "wrong format of category id", nil, nil)
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	var product models.AddProduct
	if err := c.BindJSON(&product); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	product.CategoryId = catId

	if err := validator.New().Struct(product); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "constraints not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	productResponse, err := i.inv.AddProduct(product)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to add product", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusCreated, "product added succesfully", productResponse, nil)
	c.JSON(http.StatusCreated, successRes)

}

func (i *InventoryHandler) ListProducts(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "limit number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	products, err := i.inv.ListProducts(page, limit)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve products", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully retrieved all products", products, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *InventoryHandler) AddProductImages(c *gin.Context) {
	productIdStr := c.Param("id")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "wrong format of product id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	form, err := c.MultipartForm()
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to parse mutlipartform", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	files := form.File["images"]
	if len(files) == 0 {
		errRes := response.ClientResponse(http.StatusBadRequest, "At least one image is required", nil, nil)
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	// Allowed image extensions and max size (5 MB per image)
	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
	}
	const maxFileSize = 5 * 1024 * 1024 // 5 MB

	for _, file := range files {
		// 1. Check file size
		if file.Size > maxFileSize {
			errRes := response.ClientResponse(http.StatusBadRequest, "File size exceeds limit of 5MB: "+file.Filename, nil, nil)
			c.JSON(http.StatusBadRequest, errRes)
			return
		}

		// 2. Check file extension
		ext := strings.ToLower(filepath.Ext(file.Filename))
		if !allowedExtensions[ext] {
			errRes := response.ClientResponse(http.StatusBadRequest, "Invalid file format for "+file.Filename+". Only .jpg, .jpeg, .png, .webp are allowed", nil, nil)
			c.JSON(http.StatusBadRequest, errRes)
			return
		}
	}

	urls, err := i.inv.AddProductImages(c.Request.Context(), files, productId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "Failed to upload images", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Images uploaded successfully", gin.H{"image_urls": urls}, nil)
	c.JSON(http.StatusOK, successRes)

}

func (i *InventoryHandler) GetProductByID(c *gin.Context) {
	productIdStr := c.Param("id")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "wrong format of product id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	product, err := i.inv.GetProductByID(c.Request.Context(), productId)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "failed to get product", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "successfully retrieved product details", product, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *InventoryHandler) DeleteProductImage(c *gin.Context) {
	imageIdStr := c.Param("image_id")
	imageId, err := strconv.Atoi(imageIdStr)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "wrong format of image id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := i.inv.DeleteProductImage(c.Request.Context(), imageId); err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to delete image", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "image deleted successfully from s3 and database", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

func (i *InventoryHandler) DeleteProduct(c *gin.Context) {
	productIdStr := c.Param("id")
	productId, err := strconv.Atoi(productIdStr)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "wrong format of product id", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	if err := i.inv.DeleteProduct(c.Request.Context(), productId); err != nil {
		errRes := response.ClientResponse(http.StatusInternalServerError, "failed to delete product", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "product and its images deleted successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

