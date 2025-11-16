package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/dto"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
)

type ApiTokenRepository interface {
	Create(token domain.ApiToken) error
	GetAll(userId int) []dto.ApiTokenResponse
}

type apiTokenRepository struct {
	db *sqlx.DB
}

func NewApiTokenRepository(db *sqlx.DB) ApiTokenRepository {
	return &apiTokenRepository{db: db}
}

func (r *apiTokenRepository) Create(token domain.ApiToken) error {
	_, err := r.db.NamedExec(`
		INSERT INTO api_token (user_id, token_hash, name, revoked) 
		VALUES 
		(:user_id, :token_hash, :name, :revoked)`, token)

	return err
}

func (r *apiTokenRepository) GetAll(userId int) []dto.ApiTokenResponse {
	res := []dto.ApiTokenResponse{}

	r.db.Select(&res, "SELECT id, name, revoked, created_at FROM api_token WHERE user_id = ?", userId)

	return res
}
