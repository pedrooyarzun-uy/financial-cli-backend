package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
)

type CreditCardRepository interface {
	Add(creditCard domain.CreditCard) error
	GetCurrency(creditCardId int) int
}

type creditCardRepository struct {
	db *sqlx.DB
}

func NewCreditCard(db *sqlx.DB) CreditCardRepository {
	return &creditCardRepository{db: db}
}

func (r *creditCardRepository) Add(creditCard domain.CreditCard) error {
	_, err := r.db.NamedExec(`
		INSERT INTO credit_card (name, bankID, ownerID, close_day, due_day, credit_limit)
		VALUES
		(:name, :bankID, :ownerID, :close_day, :due_day, :credit_limit)`, creditCard)

	return err
}

func (r *creditCardRepository) GetCurrency(creditCardId int) int {
	var currency int

	r.db.Get(&currency, "SELECT currency_id FROM credit_card WHERE ownerID = ?", creditCardId)

	return currency
}
