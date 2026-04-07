package dto

type AddCreditCardReq struct {
	Name        string `json:"name"`
	BankID      int    `json:"bank_id"`
	OwnerID     int
	CloseDay    int     `json:"close_day"`
	DueDay      int     `json:"due_day"`
	CreditLimit float64 `json:"limit"`
}
