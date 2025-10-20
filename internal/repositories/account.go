package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
)

type AccountRepository interface {
	Create(acc domain.Account) error
	Delete(id int) error
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

func (r *accountRepository) Delete(id int) error {
	_, err := r.db.Exec(`UPDATE account SET deleted = ? WHERE id = ?`, true, id)

	return err
}
