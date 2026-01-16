package services

import (
	"fmt"

	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/dto"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/repositories"
)

type AccountService interface {
	Create(req dto.CreateReq) error
	Delete(req dto.DeleteReq) error
	GetAll(userId int) []domain.Account
	UpdateCashBalance(acc int, amount float64, transType int) error
	GetCashBalance(acc int) float64
}

type accountService struct {
	ar repositories.AccountRepository
}

func NewAccountService(ar repositories.AccountRepository) AccountService {
	return &accountService{
		ar,
	}
}

func (s *accountService) Create(req dto.CreateReq) error {

	exists := s.ar.GetByNumber(req.Number)

	if exists.Id != 0 {
		return ErrAccountAlreadyExists
	}

	acc := domain.Account{
		Name:     req.Name,
		Number:   req.Number,
		Currency: req.Currency,
		Cash:     req.Cash,
		Bank:     req.Bank,
		Owner:    req.Owner,
		Deleted:  false,
	}

	err := s.ar.Create(acc)
	fmt.Println(err)
	return err
}

func (s *accountService) Delete(req dto.DeleteReq) error {
	return nil
}

func (s *accountService) GetAll(userId int) []domain.Account {
	return s.ar.GetAll(userId)
}

func (s *accountService) UpdateCashBalance(acc int, amount float64, transType int) error {
	return s.ar.UpdateCashBalance(acc, amount, transType)
}

func (s *accountService) GetCashBalance(acc int) float64 {
	return s.ar.GetCashBalance(acc)
}
