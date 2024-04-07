package controller

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"golang-restful-api/helper"
	"golang-restful-api/model/dto"
	"golang-restful-api/service"
	"net/http"
	"strconv"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)

	categoryCreateRequest := dto.CategoryCreateRequest{}
	err := decoder.Decode(&categoryCreateRequest)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.Create(r.Context(), categoryCreateRequest)
	webResponse := dto.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)

	categoryUpdateRequest := dto.CategoryUpdateRequest{}
	err := decoder.Decode(&categoryUpdateRequest)
	helper.PanicIfError(err)

	id, err := strconv.Atoi(params.ByName("id"))
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(r.Context(), categoryUpdateRequest)
	webResponse := dto.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	helper.PanicIfError(err)

	controller.CategoryService.Delete(r.Context(), id)
	webResponse := dto.WebResponse{
		Code:   200,
		Status: "Ok",
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(r.Context(), id)
	webResponse := dto.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryResponses := controller.CategoryService.FindAll(r.Context())
	webResponse := dto.WebResponse{
		Code:   200,
		Status: "Ok",
		Data:   categoryResponses,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}
