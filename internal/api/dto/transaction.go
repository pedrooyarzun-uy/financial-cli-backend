package dto

type AddReq struct {
	Notes       string  `json:"notes" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
	Account     int     `json:"account" binding:"required"`
	Currency    int     `json:"currenct" binding:"required"`
	Category    int     `json:"category" binding:"required"`
	Subcategory int     `json:"subcategory" binding:"required"`
	Type        int     `json:"type" binding:"required"`
}
