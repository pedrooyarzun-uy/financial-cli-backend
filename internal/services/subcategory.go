package services

import (
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/repositories"
)

type SubcategoryService interface {
	GetSubcategoriesByCategory(categoryId int, userId int) ([]domain.Subcategory, error)
}

type subcategoryService struct {
	sr repositories.SubcategoryRepository
}

func NewSubcategoryService(sr repositories.SubcategoryRepository) SubcategoryService {
	return &subcategoryService{
		sr,
	}
}

func (s *subcategoryService) GetSubcategoriesByCategory(categoryId int, userId int) ([]domain.Subcategory, error) {

	return s.sr.GetSubcategoriesByCategory(categoryId, userId)

}
