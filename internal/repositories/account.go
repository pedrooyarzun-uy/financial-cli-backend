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
	GetAll(userId int) []domain.Account
	UpdateCashBalance(acc int, amount float64, transType int) error
	GetCashBalance(acc int) float64
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

func (r *accountRepository) GetAll(userId int) []domain.Account {
	accounts := []domain.Account{}

	r.db.Select(&accounts, "SELECT * FROM account WHERE owner = ?", userId)

	return accounts
}

func (r *accountRepository) UpdateCashBalance(acc int, amount float64, transType int) error {

	var err error
	if transType == 1 {
		_, err = r.db.Exec(`UPDATE account SET balance = balance + ? WHERE id = ?`, amount, acc)
	} else {
		_, err = r.db.Exec(`UPDATE account SET balance = balance - ? WHERE id = ?`, amount, acc)
	}

	return err
}

func (r *accountRepository) GetCashBalance(acc int) float64 {
	var balance float64
	err := r.db.Get(&balance, "SELECT balance FROM account WHERE id = ?", acc)

	if err != nil {
		return 0
	}

	return balance
}
