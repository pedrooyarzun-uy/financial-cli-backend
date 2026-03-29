package services

import (
	"errors"
	"time"

	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/dto"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/repositories"
)

type TransactionService interface {
	Add(req dto.AddTransactionReq, userID int) error
	GetTotalsByCategory(userId int, from time.Time, to time.Time, category int) ([]dto.CategoryTotal, error)
	GetTransactionsByDetail(usrId int, from time.Time, to time.Time, category int, subcategory int, page int, limit int) ([]dto.TransactionByDetail, int, error)
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

func (s *transactionService) Add(req dto.AddTransactionReq, userID int) error {
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

	//TODO: Logic for currency convertion.
	if transaction.Currency != accountCurrency || accountCurrency == 0 {
		return errors.New("Transaction currency is not the same as account currency")

	}

	if transaction.Type != 1 && transaction.Type != 2 && transaction.Type != 3 {
		return errors.New("Transaction type must be 'Income', 'Outcome' or 'Adjustment'")
	}

	//---Init check belong account---
	belongsAccount, err := s.ar.BelongsToUser(transaction.Account, userID)

	if err != nil {
		return errors.New("Something went wrong")
	}

	if !belongsAccount {
		return errors.New("The selected account not belongs to you")
	}
	//---End check account---

	//--Init update cash balance---
	err = s.ar.UpdateCashBalance(transaction.Account, transaction.Amount, transaction.Type)

	if err != nil {
		return ErrCantUpdateBalance
	}
	//---End update cash balance---

	if req.Currency != accountCurrency || accountCurrency == 0 {
		return ErrTransactionNotCorrectCurrency
	}

	err = s.tr.Add(transaction)

	return err

}

func (s *transactionService) GetTotalsByCategory(userId int, from time.Time, to time.Time, category int) ([]dto.CategoryTotal, error) {

	return s.tr.GetTotalsByCategory(userId, from, to, category)
}

func (s *transactionService) GetTransactionsByDetail(usrId int, from time.Time, to time.Time, category int, subcategory int, page int, limit int) ([]dto.TransactionByDetail, int, error) {
	return s.tr.GetTransactionsByDetail(usrId, from, to, category, subcategory, page, limit)
}
