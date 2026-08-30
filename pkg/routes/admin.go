package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sangeeth518/go-Ecommerce/pkg/api/handler"
	"github.com/sangeeth518/go-Ecommerce/pkg/api/middleware"
	"github.com/sangeeth518/go-Ecommerce/pkg/config"
)

func AdminRoutes(engine *gin.RouterGroup, adminHandler *handler.AdminHandler, categoryHandler *handler.CategoryHandler, inventoryHandler *handler.InventoryHandler, cfg config.Config) {
	engine.POST("/adminlogin", adminHandler.LoginHandler)
	engine.GET("/blockuser/:id", middleware.AdminAuthMiddleware(cfg), adminHandler.BlockUser)
	engine.GET("/unblock/:id", middleware.AdminAuthMiddleware(cfg), adminHandler.UnblockUser)
	engine.GET("/getusers", middleware.AdminAuthMiddleware(cfg), adminHandler.Getusers)
	engine.POST("/addcategory", middleware.AdminAuthMiddleware(cfg), categoryHandler.AddCategory)
	engine.POST("/editcategory", middleware.AdminAuthMiddleware(cfg), categoryHandler.EditCategory)
	engine.POST("/addproduct/:category_id", middleware.AdminAuthMiddleware(cfg), inventoryHandler.AddProduct)
	engine.GET("/products", middleware.AdminAuthMiddleware(cfg), inventoryHandler.ListProducts)
	engine.GET("/product/:id", middleware.AdminAuthMiddleware(cfg), inventoryHandler.GetProductByID)
	engine.POST("/product/:id/images", middleware.AdminAuthMiddleware(cfg), inventoryHandler.AddProductImages)
	engine.DELETE("/product/image/:image_id", middleware.AdminAuthMiddleware(cfg), inventoryHandler.DeleteProductImage)
	engine.DELETE("/product/:id", middleware.AdminAuthMiddleware(cfg), inventoryHandler.DeleteProduct)
}

