package dto

import "github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"

type GetAllBankRes struct {
	Message string        `json:"message"`
	Banks   []domain.Bank `json:"banks"`
}
