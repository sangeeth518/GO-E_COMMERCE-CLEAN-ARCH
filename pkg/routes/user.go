package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sangeeth518/go-Ecommerce/pkg/api/handler"
	"github.com/sangeeth518/go-Ecommerce/pkg/api/middleware"
)

func UserRoutes(engine *gin.RouterGroup, userhandler *handler.UserHandler) {
	engine.POST("/signup", userhandler.UserSignup)
	engine.POST("/login", userhandler.Login)
	engine.PUT("/changepass", middleware.UserAuth, userhandler.ChangePassword)

}
