package handler

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"restful-api/helper"
	"restful-api/model/request"
	"restful-api/model/response"
	"restful-api/service"
	"strconv"
)

type CategoryHandlerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryHandler(CategoryService service.CategoryService) *CategoryHandlerImpl {
	return &CategoryHandlerImpl{
		CategoryService,
	}
}

func (handler *CategoryHandlerImpl) Create(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	categoryCreateRequest := request.CategoryCreateRequest{}
	helper.ReadFromRequestBody(r, &categoryCreateRequest)
	newCategory := handler.CategoryService.Create(r.Context(), categoryCreateRequest)
	apiResponse := response.Response{
		Code: http.StatusOK,
		Status: "Ok",
		Data: newCategory,
	}

	helper.ResponseHandler(w, apiResponse)
}

func (handler *CategoryHandlerImpl) Update(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	categoryUpdateRequest := request.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(r, &categoryUpdateRequest)
	categoryId := param.ByName("id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	categoryUpdateRequest.Id = id

	updateCategory := handler.CategoryService.Update(r.Context(), categoryUpdateRequest)
	apiResponse := response.Response{
		Code: http.StatusOK,
		Status: "Ok",
		Data: updateCategory,
	}

	helper.ResponseHandler(w, apiResponse)
}

func (handler *CategoryHandlerImpl) Delete(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	categoryId := param.ByName("id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	handler.CategoryService.Delete(r.Context(), id)
	apiResponse := response.Response{
		Code: http.StatusOK,
		Status: "Ok",
		Data: "",
	}

	helper.ResponseHandler(w, apiResponse)
}

func (handler *CategoryHandlerImpl) FindById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	categoryId := param.ByName("id")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	category := handler.CategoryService.FindById(r.Context(), id)
	apiResponse := response.Response{
		Code: http.StatusOK,
		Status: "Ok",
		Data: category,
	}

	helper.ResponseHandler(w, apiResponse)
}

func (handler *CategoryHandlerImpl) FindAll(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	categories := handler.CategoryService.FindAll(r.Context())
	apiResponse := response.Response{
		Code: http.StatusOK,
		Status: "Ok",
		Data: categories,
	}

	helper.ResponseHandler(w, apiResponse)
}

