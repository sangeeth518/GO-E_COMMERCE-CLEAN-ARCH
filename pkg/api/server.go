package http

import (
	"log"

	"github.com/gin-gonic/gin"
	handler "github.com/sangeeth518/go-Ecommerce/pkg/api/handler"
	"github.com/sangeeth518/go-Ecommerce/pkg/routes"
)

type ServerHttp struct {
	engine *gin.Engine
}

func NewServerHttp(adminHandler *handler.AdminHandler, userhandler *handler.UserHandler, categoryhandler *handler.CategoryHandler) *ServerHttp {

	// gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	engine.Use(gin.Logger())

	routes.AdminRoutes(engine.Group("/admin"), adminHandler, categoryhandler)
	routes.UserRoutes(engine.Group("/user"), userhandler)
	return &ServerHttp{
		engine: engine,
	}

}
func (sh *ServerHttp) Start() {
	err := sh.engine.Run(":3000")
	if err != nil {
		log.Fatal("gin engine could'nt start")
	}

}
