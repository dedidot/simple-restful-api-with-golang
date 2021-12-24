package app

import (
	"github.com/julienschmidt/httprouter"
	"restful-api/exception"
	"restful-api/handler"
)

func NewRouter(CategoryHandler handler.CategoryHandler) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/categories", CategoryHandler.FindAll)
	router.GET("/api/category/:id", CategoryHandler.FindById)
	router.POST("/api/category", CategoryHandler.Create)
	router.PUT("/api/category/:id", CategoryHandler.Update)
	router.DELETE("/api/category/:id", CategoryHandler.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}