package repositories

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
)

type AccountRepository interface {
	Create(acc domain.Account) error
	GetByNumber(number string) domain.Account
	GetCurrency(acc int) int
}

type accountRepository struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) Create(acc domain.Account) error {
	_, err := r.db.NamedExec(`
		INSERT INTO account (name, number, currency, cash, bank, owner)
		VALUES
		(:name, :number, :currency, :cash, :bank, :owner)`, acc)

	return err
}

func (r *accountRepository) GetByNumber(number string) domain.Account {
	acc := domain.Account{}

	err := r.db.Get(&acc, "SELECT * FROM account WHERE number = ? AND deleted = false", number)

	if err == sql.ErrNoRows {
		return acc
	}

	return acc

}

func (r *accountRepository) GetCurrency(acc int) int {
	var cur int

	r.db.Get(&cur, "SELECT currency FROM account WHERE id = ?", acc)

	return cur
}
