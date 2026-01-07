package dto

import "github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"

type GetSubcategoriesByCategoryRes struct {
	Message       string               `json:"message" binding:"required"`
	Subcategories []domain.Subcategory `json:"subcategories" binding:"required"`
}
