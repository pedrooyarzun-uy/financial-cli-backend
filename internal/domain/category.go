package domain

import "time"

type Category struct {
	Id        int       `db:"id"`
	Name      string    `db:"name"`
	UserId    *int      `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	Color     string    `db:"color"`
}
