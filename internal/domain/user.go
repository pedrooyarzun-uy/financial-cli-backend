package domain

import "time"

type User struct {
	Id                         int       `db:"id"`
	Name                       string    `db:"name"`
	Email                      string    `db:"email"`
	Password                   string    `db:"password"`
	VerificationToken          string    `db:"verification_token"`
	VerificationTokenExpiresAt time.Time `db:"verification_token_expires_at"`
	Verified                   bool      `db:"verified"`
	CreatedAt                  time.Time `db:"created_at"`
	Deleted                    bool      `db:"deleted"`
}
