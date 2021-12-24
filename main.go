package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"restful-api/helper"
)

func NewServer(router *httprouter.Router) *http.Server {
	return &http.Server{
		Addr:    ":2020",
		Handler: router,
	}
}

func main() {
	/*
	//db := app.NewDB()
	//validate := validator.New()
	//CategoryRepository := repository.NewCategoryRepository()
	//CategoryService := service.NewCategoryService(CategoryRepository, db, validate)
	//CategoryHandler := handler.NewCategoryHandler(CategoryService)
	//router := app.NewRouter(CategoryHandler)

	//server := http.Server{
	//	Addr:    ":2020",
	//	Handler: router,
	//}*/
	server := InitServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}