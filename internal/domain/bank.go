package domain

type Bank struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}
