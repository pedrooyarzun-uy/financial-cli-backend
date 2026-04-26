package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
)

type CreditCardRepository interface {
	Add(creditCard domain.CreditCard) error
	GetCurrency(creditCardId int) int
	BelongsToUser(creditCardID int, userID int) bool
}

type creditCardRepository struct {
	db *sqlx.DB
}

func NewCreditCard(db *sqlx.DB) CreditCardRepository {
	return &creditCardRepository{db: db}
}

func (r *creditCardRepository) Add(creditCard domain.CreditCard) error {
	_, err := r.db.NamedExec(`
		INSERT INTO credit_card (name, bankID, ownerID, close_day, due_day, credit_limit, currency_id)
		VALUES
		(:name, :bankID, :ownerID, :close_day, :due_day, :credit_limit, :currency_id)`, creditCard)

	return err
}

func (r *creditCardRepository) GetCurrency(creditCardId int) int {
	var currency int

	r.db.Get(&currency, "SELECT currency_id FROM credit_card WHERE id = ?", creditCardId)

	return currency
}

func (r *creditCardRepository) BelongsToUser(creditCardID int, userID int) bool {
	var exists bool

	err := r.db.Get(&exists, `SELECT EXISTS(
		SELECT 1
		FROM credit_card
		WHERE id = ? AND ownerID = ?
	)`, creditCardID, userID)

	if err != nil {
		return false
	}

	return exists
}
