package main

import (
	"fmt"
	"log"

	"github.com/sangeeth518/go-Ecommerce/pkg/config"
	"github.com/sangeeth518/go-Ecommerce/pkg/db"
)

func main() {

	config, configerr := config.LoadConfig()
	if configerr != nil {
		log.Fatal("cannot load config", configerr)
	}
	db, err := db.ConnectDB(config)
	if err != nil {
		fmt.Println("couldn connecttttt")
	}
	fmt.Println(db)

	fmt.Printf(config.DBHost)
}
