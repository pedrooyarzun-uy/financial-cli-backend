package services

import (
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/repositories"
)

type CurrencyService interface {
	GetAll() []domain.Currency
}

type currencyService struct {
	cr repositories.CurrencyRepository
}

func NewCurrencyService(cr repositories.CurrencyRepository) CurrencyService {
	return &currencyService{
		cr,
	}
}

func (s *currencyService) GetAll() []domain.Currency {
	response := s.cr.GetAll()

	return response
}
