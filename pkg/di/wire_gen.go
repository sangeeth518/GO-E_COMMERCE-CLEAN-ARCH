// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/sangeeth518/go-Ecommerce/pkg/api"
	"github.com/sangeeth518/go-Ecommerce/pkg/api/handler"
	"github.com/sangeeth518/go-Ecommerce/pkg/config"
	"github.com/sangeeth518/go-Ecommerce/pkg/db"
	"github.com/sangeeth518/go-Ecommerce/pkg/helper"
	"github.com/sangeeth518/go-Ecommerce/pkg/repository"
	"github.com/sangeeth518/go-Ecommerce/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHttp, error) {
	gormDB, err := db.ConnectDB(cfg)
	if err != nil {
		return nil, err
	}
	adminRepo := repository.NewAdminRepository(gormDB)
	interfacesHelper := helper.NewHelper(cfg)
	adminUseCase := usecase.NewAdminUsecase(adminRepo, interfacesHelper)
	adminHandler := handler.NewAdminHandler(adminUseCase)
	userRepo := repository.NewUserRepository(gormDB)
	userUsecase := usecase.NewUserUsecase(userRepo, interfacesHelper)
	userHandler := handler.NewUserHandler(userUsecase)
	categoryRepo := repository.NewCategoryRepository(gormDB)
	cateoryUsecase := usecase.NewCategoryUsecase(categoryRepo, interfacesHelper)
	categoryHandler := handler.NewCategoryHandler(cateoryUsecase)
	serverHttp := http.NewServerHttp(adminHandler, userHandler, categoryHandler)
	return serverHttp, nil
}
