package domain

import "time"

type ApiToken struct {
	Id        int       `db:"id"`
	UserId    int       `db:"user_id"`
	TokenHash string    `db:"token_hash"`
	Name      string    `db:"name"`
	Revoked   bool      `db:"revoked"`
	CreatedAt time.Time `db:"created_at"`
}
