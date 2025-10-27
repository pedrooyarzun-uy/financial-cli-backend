package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
)

type CurrencyRepository interface {
	GetAll() []domain.Currency
}

type currencyRepository struct {
	db *sqlx.DB
}

func NewCurrencyRepository(db *sqlx.DB) CurrencyRepository {
	return &currencyRepository{db: db}
}

func (r *currencyRepository) GetAll() []domain.Currency {
	currencies := []domain.Currency{}

	r.db.Select(&currencies, "SELECT * FROM currency ORDER BY id ASC;")

	return currencies
}
