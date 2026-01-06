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
	Color    string  `db:"color"`
}

type GetTotalsByCategoryRes struct {
	Message string          `json:"message"`
	Totals  []CategoryTotal `json:"totals"`
}

type GetCashFlowRes struct {
	Message string  `json:"message"`
	Cash    float64 `json:"cash"`
}
