package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sangeeth518/go-Ecommerce/pkg/api/handler"
	"github.com/sangeeth518/go-Ecommerce/pkg/api/middleware"
	"github.com/sangeeth518/go-Ecommerce/pkg/config"
)

func UserRoutes(engine *gin.RouterGroup, userhandler *handler.UserHandler, categoryhandler *handler.CategoryHandler, inventoryhandler *handler.InventoryHandler, carthandler *handler.CartHandler, orderhandler *handler.OrderHandler, cfg config.Config) {
	engine.POST("/signup", userhandler.UserSignup)
	engine.POST("/login", userhandler.Login)
	engine.PUT("/changepass", middleware.UserAuth(cfg), userhandler.ChangePassword)
	engine.POST("/adress", middleware.UserAuth(cfg), userhandler.AddAdress)
	engine.GET("/showcategories", middleware.UserAuth(cfg), categoryhandler.ShowCategories)
	engine.GET("/products", middleware.UserAuth(cfg), inventoryhandler.ListProducts)
	engine.GET("/product/:id", middleware.UserAuth(cfg), inventoryhandler.GetProductByID)

	// Cart routes
	engine.POST("/cart", middleware.UserAuth(cfg), carthandler.AddToCart)
	engine.GET("/cart", middleware.UserAuth(cfg), carthandler.ViewCart)
	engine.PUT("/cart", middleware.UserAuth(cfg), carthandler.UpdateQuantity)
	engine.DELETE("/cart/:id", middleware.UserAuth(cfg), carthandler.RemoveProductFromCart)

	// Order routes
	engine.POST("/order", middleware.UserAuth(cfg), orderhandler.OrderCheckout)
	engine.GET("/order", middleware.UserAuth(cfg), orderhandler.GetMyOrder)
	engine.GET("/order/:id", middleware.UserAuth(cfg), orderhandler.GetOrderDetails)
	engine.GET("/order/:id/invoice", middleware.UserAuth(cfg), orderhandler.DownloadInvoice)
	engine.PATCH("/order/:id/cancel", middleware.UserAuth(cfg), orderhandler.CancelOrder)
}
