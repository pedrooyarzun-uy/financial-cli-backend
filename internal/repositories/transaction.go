package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/dto"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
)

type TransactionRepository interface {
	Add(transaction domain.Transaction) error
	GetTotalsByCategory() []dto.CategoryTotal
	GetCashFlow() float64
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

func (r *transactionRepository) GetTotalsByCategory() []dto.CategoryTotal {
	res := []dto.CategoryTotal{}

	r.db.Select(&res, `
		select 
			c.name, 
			SUM(
				CASE
					WHEN t.type = 1 THEN t.amount
					ELSE -t.amount
				END
			) AS total
		from transaction t
		left join category c on c.id = t.category 
		group by t.category, c.name
	`)

	return res
}

func (r *transactionRepository) GetCashFlow() float64 {

	var cash float64

	r.db.Get(&cash, `
		select
			SUM(
				CASE 
					WHEN t.type = 1 THEN t.amount
					ELSE -t.amount
				END
			) AS total
		from transaction t
	`)

	return cash
}
