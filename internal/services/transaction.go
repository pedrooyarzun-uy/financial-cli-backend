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
	tr  repositories.TransactionRepository
	ar  repositories.AccountRepository
	ccr repositories.CreditCardRepository
}

func NewTransactionRepository(tr repositories.TransactionRepository, ar repositories.AccountRepository, ccr repositories.CreditCardRepository) TransactionService {
	return &transactionService{
		tr, ar, ccr,
	}
}

func (s *transactionService) Add(req dto.AddTransactionReq, userID int) error {
	transaction := domain.Transaction{
		Notes:         req.Notes,
		Amount:        req.Amount,
		Kind:          domain.TransactionKind(req.Kind),
		PaymentMethod: domain.PaymentMethod(req.PaymentMethod),
		CurrencyId:    req.CurrencyId,
		CategoryId:    req.CategoryId,
		SubcategoryId: req.SubcategoryId,
		AccountId:     req.AccountId,
		CreditCardId:  req.CreditCardId,
	}

	if req.AccountId == 0 && req.CreditCardId == 0 {
		return errors.New("You must select an account or a credit card")
	}

	if req.AccountId != 0 && req.CreditCardId != 0 {
		return errors.New("You can only select or an account or a credit card")
	}

	var currency int
	var belongs bool

	if req.AccountId != 0 {
		currency = s.ar.GetCurrency(req.AccountId)
		belongs = s.ar.BelongsToUser(transaction.AccountId, userID)
	} else {
		currency = s.ccr.GetCurrency(req.CreditCardId)
		belongs = s.ccr.BelongsToUser(req.CreditCardId, userID)
	}

	//TODO: Logic for currency convertion.
	if transaction.CurrencyId != currency || currency == 0 {
		return errors.New("Transaction currency is not the same as account currency")
	}

	if !belongs {
		return errors.New("The selected account not belongs to you")
	}

	if transaction.Kind != domain.TransactionKindExpense && transaction.Kind != domain.TransactionKindIncome {
		return errors.New("Transaction type must be 'Income', 'Outcome' or 'Adjustment'")
	}

	//TODO: Strategy pattern for solving cash balance between credit_cards and transactions
	if req.AccountId != 0 {
		err := s.ar.UpdateCashBalance(transaction.AccountId, transaction.Amount, transaction.Kind)

		if err != nil {
			return ErrCantUpdateBalance
		}
	}
	//---End update cash balance---
	err := s.tr.Add(transaction)

	return err
}

func (s *transactionService) GetTotalsByCategory(userId int, from time.Time, to time.Time, category int) ([]dto.CategoryTotal, error) {

	return s.tr.GetTotalsByCategory(userId, from, to, category)
}

func (s *transactionService) GetTransactionsByDetail(usrId int, from time.Time, to time.Time, category int, subcategory int, page int, limit int) ([]dto.TransactionByDetail, int, error) {
	return s.tr.GetTransactionsByDetail(usrId, from, to, category, subcategory, page, limit)
}
