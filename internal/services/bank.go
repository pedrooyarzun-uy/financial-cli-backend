package services

import (
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/repositories"
)

type BankService interface {
	GetAll() []domain.Bank
}

type bankService struct {
	br repositories.BankRepository
}

func NewBankService(br repositories.BankRepository) BankService {
	return &bankService{br}
}

func (s *bankService) GetAll() []domain.Bank {
	banks := s.br.GetAll()

	return banks
}
