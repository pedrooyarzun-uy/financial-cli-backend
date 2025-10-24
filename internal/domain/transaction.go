package domain

type Transaction struct {
	Id          int     `db:"id"`
	Notes       string  `db:"notes"`
	Amount      float64 `db:"amount"`
	Account     int     `db:"account"`
	Currency    int     `db:"currency"`
	Category    int     `db:"category"`
	Subcategory int     `db:"subcategory"`
	Type        int     `db:"type"`
}
