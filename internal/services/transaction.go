package services

import (
	"fmt"

	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/dto"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/repositories"
)

type TransactionService interface {
	Add(req dto.AddTransactionReq) error
	GetTotalsByCategory(userId int) []dto.CategoryTotal
	GetCashFlow(userId int) float64
}

type transactionService struct {
	tr repositories.TransactionRepository
	ar repositories.AccountRepository
}

func NewTransactionRepository(tr repositories.TransactionRepository, ar repositories.AccountRepository) TransactionService {
	return &transactionService{
		tr, ar,
	}
}

func (s *transactionService) Add(req dto.AddTransactionReq) error {
	transaction := domain.Transaction{
		Notes:       req.Notes,
		Amount:      req.Amount,
		Account:     req.Account,
		Currency:    req.Currency,
		Category:    req.Category,
		Subcategory: req.Subcategory,
		Type:        req.Type,
	}

	accountCurrency := s.ar.GetCurrency(req.Account)

	if req.Currency != accountCurrency || accountCurrency == 0 {
		fmt.Println("Currency de req.Currency: ", req.Currency)
		fmt.Println("Currency de account: ", accountCurrency)
		return ErrTransactionNotCorrectCurrency
	}

	err := s.tr.Add(transaction)

	return err

}

func (s *transactionService) GetTotalsByCategory(userId int) []dto.CategoryTotal {

	return s.tr.GetTotalsByCategory(userId)
}

func (s *transactionService) GetCashFlow(userId int) float64 {
	return s.tr.GetCashFlow(userId)
}
