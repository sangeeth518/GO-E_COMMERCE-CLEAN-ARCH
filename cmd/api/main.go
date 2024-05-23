package main

import (
	"log"

	"github.com/sangeeth518/go-Ecommerce/pkg/config"
	"github.com/sangeeth518/go-Ecommerce/pkg/di"
)

func main() {

	config, configerr := config.LoadConfig()
	if configerr != nil {
		log.Fatal("cannot load config", configerr)
	}

	server, dierr := di.InitializeAPI(config)

	if dierr != nil {
		log.Fatal("cannot start server", dierr)
	} else {
		server.Start()
	}

}
