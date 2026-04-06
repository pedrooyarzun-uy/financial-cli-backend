package domain

import "time"

type CreditCard struct {
	Id        int       `db:"id"`
	Name      string    `db:"name"`
	BankID    int       `db:"bankID"`
	OwnerID   int       `db:"ownerID"`
	CloseDay  time.Time `db:"close_day"`
	DueDay    time.Time `db:"due_day"`
	Limit     float64   `db:"limit"`
	CreatedAt time.Time `db:"created_at"`
}
