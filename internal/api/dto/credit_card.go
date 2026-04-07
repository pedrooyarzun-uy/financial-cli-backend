package dto

type AddCreditCardReq struct {
	Name     string  `json:"name"`
	BankID   int     `json:"bankID"`
	OwnerID  int     `json:"ownerID"`
	CloseDay int     `json:"close_day"`
	DueDay   int     `json:"due_day"`
	Limit    float64 `json:"limit"`
}
