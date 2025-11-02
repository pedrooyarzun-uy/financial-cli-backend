package services

import (
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/repositories"
)

type CategoryService interface {
	GetAll() []domain.Category
}

type categoryService struct {
	cr repositories.CategoryRepository
}

func NewCategoryService(cr repositories.CategoryRepository) CategoryService {
	return &categoryService{
		cr,
	}
}

func (s *categoryService) GetAll() []domain.Category {
	return s.cr.GetAll()
}
