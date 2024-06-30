//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "github.com/sangeeth518/go-Ecommerce/pkg/api"
	handler "github.com/sangeeth518/go-Ecommerce/pkg/api/handler"
	"github.com/sangeeth518/go-Ecommerce/pkg/config"
	"github.com/sangeeth518/go-Ecommerce/pkg/db"
	"github.com/sangeeth518/go-Ecommerce/pkg/helper"
	"github.com/sangeeth518/go-Ecommerce/pkg/repository"
	"github.com/sangeeth518/go-Ecommerce/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*http.ServerHttp, error) {
	wire.Build(db.ConnectDB, helper.NewHelper, repository.NewAdminRepository, usecase.NewAdminUsecase, handler.NewAdminHandler,
		repository.NewUserRepository, usecase.NewUserUsecase, handler.NewUserHandler,
		repository.NewCategoryRepository, usecase.NewCategoryUsecase, handler.NewCategoryHandler,

		http.NewServerHttp)

	return &http.ServerHttp{}, nil
}
