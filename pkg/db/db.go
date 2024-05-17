package db

import (
	"fmt"

	"github.com/sangeeth518/go-Ecommerce/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPassword, cfg.DBPort)
	db, dberr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{SkipDefaultTransaction: true})
	if dberr != nil {
		fmt.Println("couldn't connect db")
	}
	db.AutoMigrate()
	return db, dberr
}
