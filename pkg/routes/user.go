package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sangeeth518/go-Ecommerce/pkg/api/handler"
	"github.com/sangeeth518/go-Ecommerce/pkg/api/middleware"
	"github.com/sangeeth518/go-Ecommerce/pkg/config"
)

func UserRoutes(engine *gin.RouterGroup, userhandler *handler.UserHandler, categoryhandler *handler.CategoryHandler, cfg config.Config) {
	engine.POST("/signup", userhandler.UserSignup)
	engine.POST("/login", userhandler.Login)
	engine.PUT("/changepass", middleware.UserAuth(cfg), userhandler.ChangePassword)
	engine.POST("/adress", middleware.UserAuth(cfg), userhandler.AddAdress)
	engine.GET("/showcategories", middleware.UserAuth(cfg), categoryhandler.ShowCategories)

}
