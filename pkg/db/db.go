package db

import (
	"fmt"

	"github.com/sangeeth518/go-Ecommerce/pkg/config"
	"github.com/sangeeth518/go-Ecommerce/pkg/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPassword, cfg.DBPort)
	db, dberr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{SkipDefaultTransaction: true})
	if dberr != nil {
		fmt.Println("couldn't connect db")
	}
	db.AutoMigrate(&domain.Admin{})
	CheckAndCreateAdmin(db)
	return db, dberr
}

func CheckAndCreateAdmin(db *gorm.DB) {
	var count int64
	db.Model(&domain.Admin{}).Count(&count)
	if count == 0 {
		password := "adminpass"
		hashpass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
		if err != nil {
			fmt.Println("admin pass hashing err")
			return
		}
		admin := domain.Admin{
			Id:       1,
			Name:     "Sarang",
			Email:    "sarangsajeev@gmail.com",
			Password: string(hashpass),
		}
		db.Create(&admin)

	}
}
