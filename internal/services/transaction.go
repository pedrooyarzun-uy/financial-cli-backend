package services

import (
	"fmt"
	"time"

	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/dto"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/repositories"
)

type TransactionService interface {
	Add(req dto.AddTransactionReq) error
	GetTotalsByCategory(userId int) []dto.CategoryTotal
	GetCashFlow(userId int) float64
	GetTransactionsByDetail(usrId int, from time.Time, to time.Time, category int, subcategory int) ([]dto.TransactionByDetail, error)
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

	err := s.ar.UpdateCashBalance(transaction.Account, transaction.Amount, transaction.Type)

	if err != nil {
		return ErrCantUpdateBalance
	}

	if req.Currency != accountCurrency || accountCurrency == 0 {
		fmt.Println("Currency de req.Currency: ", req.Currency)
		fmt.Println("Currency de account: ", accountCurrency)
		return ErrTransactionNotCorrectCurrency
	}

	err = s.tr.Add(transaction)

	return err

}

func (s *transactionService) GetTotalsByCategory(userId int) []dto.CategoryTotal {

	return s.tr.GetTotalsByCategory(userId)
}

func (s *transactionService) GetCashFlow(userId int) float64 {
	return s.tr.GetCashFlow(userId)
}

func (s *transactionService) GetTransactionsByDetail(usrId int, from time.Time, to time.Time, category int, subcategory int) ([]dto.TransactionByDetail, error) {
	return s.tr.GetTransactionsByDetail(usrId, from, to, category, subcategory)
}
