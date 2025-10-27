package domain

type Currency struct {
	Id     int    `db:"id"`
	Name   string `db:"name"`
	Symbol rune   `db:"symbol"`
	Code   string `db:"code"`
}
