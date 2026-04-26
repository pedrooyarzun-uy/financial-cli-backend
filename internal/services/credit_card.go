package services

import (
	"errors"

	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/dto"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/repositories"
)

type CreditCardService interface {
	Add(req dto.AddCreditCardReq) error
}

type creditCardService struct {
	ccr repositories.CreditCardRepository
}

func NewCreditCardService(ccr repositories.CreditCardRepository) CreditCardService {
	return &creditCardService{ccr}
}

func (s *creditCardService) Add(req dto.AddCreditCardReq) error {
	if req.Name == "" {
		return errors.New("Name can't be empty")
	}

	if req.BankID == 0 {
		return errors.New("Bank can't be empty")
	}

	if req.CloseDay < 1 || req.CloseDay > 31 {
		return errors.New("Close day must be in the range between 1 and 31")
	}

	if req.DueDay < 1 || req.CloseDay > 31 {
		return errors.New("Due day must be in the range between 1 and 31")
	}

	if req.CreditLimit == 0.00 {
		return errors.New("Limit can't be empty")
	}

	if req.CurrencyId == 0 {
		return errors.New("Currency can't be empty")
	}

	creditCard := domain.CreditCard{
		Name:        req.Name,
		BankID:      req.BankID,
		CloseDay:    req.CloseDay,
		DueDay:      req.DueDay,
		CreditLimit: req.CreditLimit,
		OwnerID:     req.OwnerID,
		CurrencyId:  req.CurrencyId,
	}

	err := s.ccr.Add(creditCard)

	if err != nil {
		return errors.New("Something went wrong while trying to add your account")
	}

	return nil
}
