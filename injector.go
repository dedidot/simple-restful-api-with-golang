//go:build wireinject
//+build wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"net/http"
	"restful-api/app"
	"restful-api/handler"
	"restful-api/repository"
	"restful-api/service"
)

var CategorySet = wire.NewSet(
		repository.NewCategoryRepository,
		wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
		service.NewCategoryService,
		wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
		handler.NewCategoryHandler,
		wire.Bind(new(handler.CategoryHandler), new(*handler.CategoryHandlerImpl)),
	)

func InitServer() *http.Server {
	wire.Build(
			app.NewDB,
			validator.New,
			CategorySet,
			app.NewRouter,
			//wire.Bind(new(http.Handler), new(*httprouter.Router)),
			//app.NewRouter,
			NewServer,
		)
	return nil
}