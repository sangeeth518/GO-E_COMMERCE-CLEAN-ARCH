package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sangeeth518/go-Ecommerce/pkg/api/handler"
)

func UserRoutes(engine *gin.RouterGroup, userhandler *handler.UserHandler) {
	engine.POST("/signup", userhandler.UserSignup)

}
