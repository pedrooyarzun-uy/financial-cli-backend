package dto

type AddTransactionReq struct {
	Notes       string  `json:"notes"`
	Amount      float64 `json:"amount" binding:"required"`
	Account     int     `json:"account" binding:"required"`
	Currency    int     `json:"currency" binding:"required"`
	Category    int     `json:"category"`
	Subcategory int     `json:"subcategory"`
	Type        int     `json:"type" binding:"required"`
}

type CategoryTotal struct {
	Category string  `db:"name"`
	Total    float64 `db:"total"`
}

type GetTotalsByCategoryRes struct {
	Message string          `json:"message"`
	Totals  []CategoryTotal `json:"totals"`
}
