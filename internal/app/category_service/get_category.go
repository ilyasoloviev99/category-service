package category_service

import (
	"context"
	"fmt"
	desc "github.com/ilyasoloviev99/category-service/pkg/category"
)

type CategoryService struct {
	desc.UnimplementedCategoryServiceServer
}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}

func (s *CategoryService) GetCategory(
	ctx context.Context, req *desc.GetCategoryRequest,
) (*desc.GetCategoryResponse, error) {
	name := req.GetName()

	return &desc.GetCategoryResponse{
		Name: fmt.Sprintf("Hello %s", name),
	}, nil
}
