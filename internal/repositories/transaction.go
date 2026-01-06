package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/dto"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
)

type TransactionRepository interface {
	Add(transaction domain.Transaction) error
	GetTotalsByCategory(userId int) []dto.CategoryTotal
	GetCashFlow(userId int) float64
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

func (r *transactionRepository) GetTotalsByCategory(userId int) []dto.CategoryTotal {
	res := []dto.CategoryTotal{}

	r.db.Select(&res, `
		select 
			c.name, 
			SUM(t.amount) AS total,
			c.color
		from transaction t
		left join category c on c.id = t.category 
		join account a on a.id = t.account
		where a.owner = ? and t.type = 2
		group by t.category, c.name
	`, userId)

	return res
}

func (r *transactionRepository) GetCashFlow(userId int) float64 {

	var cash float64

	r.db.Get(&cash, `
		select
			SUM(
				CASE 
					WHEN t.type = 1 THEN -t.amount
					ELSE t.amount
				END
			) AS total
		from transaction t
		join account a on a.id = t.account
		where a.owner = ?
	`, userId)

	return cash
}
