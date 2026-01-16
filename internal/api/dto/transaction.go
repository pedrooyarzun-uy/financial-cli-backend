package dto

import (
	"time"
)

type AddTransactionReq struct {
	Notes       string  `json:"notes"`
	Amount      float64 `json:"amount" binding:"required"`
	Account     int     `json:"account" binding:"required"`
	Currency    int     `json:"currency" binding:"required"`
	Category    int     `json:"category"`
	Subcategory int     `json:"subcategory"`
	Type        int     `json:"type" binding:"required"`
}

// Begin Method
type CategoryTotal struct {
	Category string  `db:"name"`
	Total    float64 `db:"total"`
	Color    string  `db:"color"`
}

type GetTotalsByCategoryRes struct {
	Message string          `json:"message"`
	Totals  []CategoryTotal `json:"totals"`
}

//End Method

type GetCashFlowRes struct {
	Message string  `json:"message"`
	Cash    float64 `json:"cash"`
}

type TransactionByDetail struct {
	Id          int       `db:"id"`
	Category    string    `db:"category"`
	Subcategory string    `db:"subcategory"`
	Amount      float64   `db:"amount"`
	Currency    string    `db:"currency"`
	Notes       string    `db:"notes"`
	Date        time.Time `db:"created_at"`
	Color       string    `db:"color"`
}

type GetTransactionsByDetailRes struct {
	Message      string                `json:"message"`
	Transactions []TransactionByDetail `json:"transactions"`
}
