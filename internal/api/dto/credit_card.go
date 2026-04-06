package dto

import "time"

type AddCreditCardReq struct {
	Name     string    `json:"name"`
	BankID   int       `json:"bankID"`
	OwnerID  int       `json:"ownerID"`
	CloseDay time.Time `json:"close_day"`
	DueDay   time.Time `json:"due_day"`
	Limit    float64   `json:"limit"`
}
