package helper

import (
	"golang-restful-api/model/domain"
	"golang-restful-api/model/dto"
)

func ToCategoryResponse(category domain.Category) dto.CategoryResponse {
	return dto.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
