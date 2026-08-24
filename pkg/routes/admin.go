package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sangeeth518/go-Ecommerce/pkg/api/handler"
	"github.com/sangeeth518/go-Ecommerce/pkg/api/middleware"
	"github.com/sangeeth518/go-Ecommerce/pkg/config"
)

func AdminRoutes(engine *gin.RouterGroup, adminHandler *handler.AdminHandler, categoryHandler *handler.CategoryHandler, cfg config.Config) {
	engine.POST("/adminlogin", adminHandler.LoginHandler)
	engine.GET("/blockuser/:id", middleware.AdminAuthMiddleware(cfg), adminHandler.BlockUser)
	engine.GET("/unblock/:id", middleware.AdminAuthMiddleware(cfg), adminHandler.UnblockUser)
	engine.GET("/getusers", middleware.AdminAuthMiddleware(cfg), adminHandler.Getusers)
	engine.POST("/addcategory", middleware.AdminAuthMiddleware(cfg), categoryHandler.AddCategory)
	engine.POST("/editcategory", middleware.AdminAuthMiddleware(cfg), categoryHandler.EditCategory)
}
