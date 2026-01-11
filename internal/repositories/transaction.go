package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/dto"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
)

type TransactionRepository interface {
	Add(transaction domain.Transaction) error
	GetTotalsByCategory(userId int) []dto.CategoryTotal
	GetCashFlow(userId int) float64
	GetTransactionsByDetail(usrId int, from time.Time, to time.Time, category int, subcategory int) ([]dto.TransactionByDetail, error)
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

func (r *transactionRepository) GetTransactionsByDetail(usrId int, from time.Time, to time.Time, category int, subcategory int) ([]dto.TransactionByDetail, error) {

	var res []dto.TransactionByDetail

	query := `
		SELECT 
			t.id, 
			COALESCE(c.name, 'N/A') AS 'category', 
			COALESCE(s.name, 'N/A') AS 'subcategory', 
			t.amount, 
			COALESCE(t.notes, 'N/A') AS 'notes',
			t.created_at,
			cu.symbol AS 'currency'
		FROM transaction t 
		LEFT JOIN category c ON c.id = t.category
		LEFT JOIN subcategory s ON t.subcategory = s.id
		JOIN account a ON a.id = t.account
		JOIN currency cu ON cu.id = t.currency
		WHERE a.owner = ?
	`
	args := []any{usrId}

	if category != 0 {
		query += ` AND t.category = ?`
		args = append(args, category)
	}

	if subcategory != 0 {
		query += ` AND t.subcategory = ?`
		args = append(args, subcategory)
	}

	now := time.Now()

	if from.IsZero() || to.IsZero() {
		start := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
		end := start.AddDate(0, 1, 0)

		query += " AND t.created_at >= ? AND t.created_at < ?"
		args = append(args, start, end)

	} else {
		if to.Sub(from) > time.Hour*24*30*6 {
			start := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
			end := start.AddDate(0, 1, 0)

			query += " AND t.created_at >= ? AND t.created_at < ?"
			args = append(args, start, end)
		} else {
			query += " AND t.created_at BETWEEN ? AND ?"
			args = append(args, from, to)
		}
	}

	query += " ORDER BY t.created_at DESC;"

	err := r.db.Select(&res, query, args...)

	return res, err

}
