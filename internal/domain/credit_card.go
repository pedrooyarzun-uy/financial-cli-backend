package domain

import "time"

type CreditCard struct {
	Id        int       `db:"id"`
	Name      string    `db:"name"`
	BankID    int       `db:"bankID"`
	OwnerID   int       `db:"ownerID"`
	CloseDay  int       `db:"close_day"`
	DueDay    int       `db:"due_day"`
	Limit     float64   `db:"limit"`
	CreatedAt time.Time `db:"created_at"`
}
