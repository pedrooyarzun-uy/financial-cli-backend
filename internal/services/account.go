package services

import (
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/dto"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/repositories"
)

type AccountService interface {
	Create(req dto.CreateReq) error
	Delete(req dto.DeleteReq) error
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

	return err
}

func (s *accountService) Delete(req dto.DeleteReq) error {
	return nil
}
