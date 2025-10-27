package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
)

type BankRepository interface {
	GetAll() []domain.Bank
}

type bankRepository struct {
	db *sqlx.DB
}

func NewBankRepository(db *sqlx.DB) BankRepository {
	return &bankRepository{db: db}
}

func (r *bankRepository) GetAll() []domain.Bank {
	banks := []domain.Bank{}

	r.db.Select(&banks, "SELECT * FROM bank ORDER BY name ASC")

	return banks
}
