package dto

import "github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"

type GetAllCurrencyRes struct {
	Message    string
	Currencies []domain.Currency
}
