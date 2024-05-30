package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sangeeth518/go-Ecommerce/pkg/api/handler"
)

func AdminRoutes(engine *gin.RouterGroup, adminHandler *handler.AdminHandler) {
	engine.POST("/adminlogin", adminHandler.LoginHandler)
	engine.GET("/blockuser/:id", adminHandler.BlockUser)
	engine.GET("/unblock/:id", adminHandler.UnblockUser)
	engine.GET("/getusers", adminHandler.Getusers)
}
