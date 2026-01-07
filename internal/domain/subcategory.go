package domain

type Subcategory struct {
	Id         int    `db:"id"`
	Name       string `db:"name"`
	CategoryId int    `db:"category_id"`
	UserId     *int   `db:"user_id"`
}
