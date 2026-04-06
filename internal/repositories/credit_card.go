package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
)

type CreditCardRepository interface {
	Add(creditCard domain.CreditCard) error
}

type creditCardRepository struct {
	db *sqlx.DB
}

func NewCreditCard(db *sqlx.DB) CreditCardRepository {
	return &creditCardRepository{db: db}
}

func (r *creditCardRepository) Add(creditCard domain.CreditCard) error {
	_, err := r.db.NamedExec(`
		INSERT INTO credit_card (name, bankID, currencyID, ownerID, close_day, due_day, limit)
		VALUES
		(:name, :bankID, :currencyID, :ownerID, :close_day, :due_day, :limit)`, creditCard)

	return err
}
