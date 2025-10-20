package domain

type Account struct {
	Id       int     `db:"id"`
	Name     string  `db:"name"`
	Number   string  `db:"number"`
	Currency int     `db:"currency"`
	Cash     float64 `db:"cash"`
	Bank     int     `db:"bank"`
	Owner    int     `db:"owner"`
	Deleted  bool    `db:"deleted"`
}
