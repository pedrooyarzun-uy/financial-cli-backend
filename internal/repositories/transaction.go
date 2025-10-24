package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
)

type TransactionRepository interface {
	Add(transaction domain.Transaction) error
}

type transactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Add(transaction domain.Transaction) error {
	_, err := r.db.NamedExec(`
		INSERT INTO transaction (notes, amount, account, currency, category, subcategory, type)
		VALUES
		(:notes, :amount, :account, :currency, :category, :subcategory, :type)`, transaction)

	return err
}
