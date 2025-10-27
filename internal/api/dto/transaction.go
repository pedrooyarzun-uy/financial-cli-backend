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
