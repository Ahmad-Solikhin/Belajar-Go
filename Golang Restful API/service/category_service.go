package service

import (
	"context"
	"golang-restful-api/model/dto"
)

type CategoryService interface {
	Create(ctx context.Context, request dto.CategoryCreateRequest) dto.CategoryResponse
	Update(ctx context.Context, request dto.CategoryUpdateRequest) dto.CategoryResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) dto.CategoryResponse
	FindAll(ctx context.Context) []dto.CategoryResponse
}
