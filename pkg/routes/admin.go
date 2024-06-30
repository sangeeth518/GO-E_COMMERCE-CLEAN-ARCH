package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sangeeth518/go-Ecommerce/pkg/api/handler"
	"github.com/sangeeth518/go-Ecommerce/pkg/api/middleware"
)

func AdminRoutes(engine *gin.RouterGroup, adminHandler *handler.AdminHandler, categoryHandler *handler.CategoryHandler) {
	engine.POST("/adminlogin", adminHandler.LoginHandler)
	engine.GET("/blockuser/:id", middleware.AdminAuthMiddleware, adminHandler.BlockUser)
	engine.GET("/unblock/:id", middleware.AdminAuthMiddleware, adminHandler.UnblockUser)
	engine.GET("/getusers", middleware.AdminAuthMiddleware, adminHandler.Getusers)
	engine.POST("/addcategory", middleware.AdminAuthMiddleware, categoryHandler.AddCategory)
}
