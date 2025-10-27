package domain

type Currency struct {
	Id     int    `db:"id"`
	Name   string `db:"name"`
	Symbol string `db:"symbol"`
	Code   string `db:"code"`
}
