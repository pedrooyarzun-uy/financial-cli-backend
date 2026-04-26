package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/api/dto"
	"github.com/pedrooyarzun-uy/financial-cli-backend/internal/domain"
)

type TransactionRepository interface {
	Add(transaction domain.Transaction) error
	GetTotalsByCategory(userId int, from time.Time, to time.Time, category int) ([]dto.CategoryTotal, error)
	GetTransactionsByDetail(usrId int, from time.Time, to time.Time, category int, subcategory int, page int, limit int) ([]dto.TransactionByDetail, int, error)
}

type transactionRepository struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Add(transaction domain.Transaction) error {
	_, err := r.db.NamedExec(`
		INSERT INTO transaction (notes, amount, kind, payment_method, currency_id, category_id, subcategory_id, account_id, credit_card_id)
		VALUES
		(:notes, :amount, :kind, :payment_method, :currency_id, NULLIF(:category_id, 0), NULLIF(:subcategory_id, 0), NULLIF(:account_id, 0), NULLIF(:credit_card_id, 0))`, transaction)

	return err
}

func (r *transactionRepository) GetTotalsByCategory(userId int, from time.Time, to time.Time, category int) ([]dto.CategoryTotal, error) {

	query := ""

	if category != 0 {
		query += `SELECT s.name, `
	} else {
		query += `SELECT c.name, `
	}

	query += `SUM(t.amount) AS total, c.color
		FROM transaction t
		LEFT JOIN category c ON c.id = t.category 
		LEFT JOIN subcategory s ON s.id = t.subcategory
		JOIN account a ON a.id = t.account
		WHERE a.owner = ? 
	`

	args := []any{userId}

	if category != 0 {
		query += ` AND t.category = ?`
		args = append(args, category)
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

	query += " AND t.type = 2"

	if category != 0 {
		query += " GROUP BY t.subcategory, s.name, c.color"
	} else {
		query += " GROUP BY t.category, c.name, c.color"
	}

	res := []dto.CategoryTotal{}

	err := r.db.Select(&res, query, args...)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *transactionRepository) GetTransactionsByDetail(usrId int, from time.Time, to time.Time, category int, subcategory int, page int, limit int) ([]dto.TransactionByDetail, int, error) {

	var res []dto.TransactionByDetail

	query := `
		SELECT 
			t.id, 
			COALESCE(c.name, IF(t.type = 3, 'Adjustment', 'N/A')) AS 'category', 
			COALESCE(s.name, IF(t.type = 3, 'Adjustment', 'N/A')) AS 'subcategory', 
			t.amount, 
			COALESCE(t.notes, 'N/A') AS 'notes',
			t.created_at,
			cu.symbol AS 'currency',
			COALESCE(c.color, '') AS color,
			t.type
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

	query += " ORDER BY t.created_at DESC"

	results := []dto.TransactionByDetail{}
	err := r.db.Select(&results, query, args...)

	//Get for adding
	totalResults := len(results)
	maxPage := (totalResults + limit - 1) / limit

	//Pagination
	offset := (page - 1) * limit
	query += " LIMIT ? OFFSET ?;"
	args = append(args, limit, offset)

	err = r.db.Select(&res, query, args...)

	return res, maxPage, err

}
