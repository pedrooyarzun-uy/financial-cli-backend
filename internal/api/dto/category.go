package dto

import "github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"

type GetAllCategoryRes struct {
	Message    string `json:"message"`
	Categories []domain.Category
}
