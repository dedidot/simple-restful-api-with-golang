package exception

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"restful-api/helper"
	"restful-api/model/response"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if notFound(w, r, err) {
		return
	}

	if validatorError(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func validatorError(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		res := response.Response{
			Code: http.StatusBadRequest,
			Status: "Error",
			Data: exception.Error(),
		}

		helper.ResponseHandler(w, res)
		return true
	} else {
		return false
	}
}

func notFound(w http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		res := response.Response{
			Code: http.StatusNotFound,
			Status: "Error",
			Data: exception.Error,
		}

		helper.ResponseHandler(w, res)
		return true
	} else {
		return false
	}
}

func internalServerError(w http.ResponseWriter, r *http.Request, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	res := response.Response{
		Code: http.StatusInternalServerError,
		Status: "Error",
		Data: nil,
	}

	helper.ResponseHandler(w, res)
}
