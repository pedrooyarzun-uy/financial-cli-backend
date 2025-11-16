package services

import (
	"errors"

	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/dto"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/helpers"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/repositories"
)

type ApiTokenService interface {
	Create(userId int, req dto.CreateApiTokenReq) error
	GetAll(userId int) []dto.ApiTokenResponse
}

type apiTokenService struct {
	atr repositories.ApiTokenRepository
}

func NewApiTokenService(atr repositories.ApiTokenRepository) ApiTokenService {
	return &apiTokenService{atr}
}

func (s *apiTokenService) Create(userId int, req dto.CreateApiTokenReq) error {

	hash, _ := helpers.HashPassword(req.Keyword)

	token := domain.ApiToken{
		UserId:    userId,
		TokenHash: hash,
		Name:      req.Name,
		Revoked:   false,
	}

	err := s.atr.Create(token)

	if err != nil {
		return errors.New("something went wrong while saving your token")
	}

	return err
}

func (s *apiTokenService) GetAll(userId int) []dto.ApiTokenResponse {
	tokens := s.atr.GetAll(userId)

	return tokens
}
