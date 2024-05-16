package main

import (
	"fmt"
	"log"

	"github.com/sangeeth518/go-Ecommerce/pkg/config"
)

func main() {

	config, configerr := config.LoadConfig()
	if configerr != nil {
		log.Fatal("cannot load config", configerr)
	}

	fmt.Printf(config.DBHost)
}
