package domain

import "time"

type Transaction struct {
	Id            int             `db:"id"`
	Notes         string          `db:"notes"`
	Amount        float64         `db:"amount"`
	Kind          TransactionKind `db:"kind"`
	PaymentMethod PaymentMethod   `db:"payment_method"`
	CurrencyId    int             `db:"currency_id"`
	CategoryId    int             `db:"category_id"`
	SubcategoryId int             `db:"subcategory_id"`
	AccountId     int             `db:"account_id"`
	CreditCardId  int             `db:"credit_card_id"`
	CreatedAt     time.Time       `db:"created_at"`
}

type PaymentMethod string

const (
	PaymentMethodAccount    PaymentMethod = "account"
	PaymentMethodCreditCard PaymentMethod = "credit_card"
	PaymentMethodCash       PaymentMethod = "cash"
)

type TransactionKind string

const (
	TransactionKindIncome  TransactionKind = "income"
	TransactionKindExpense TransactionKind = "expense"
)
