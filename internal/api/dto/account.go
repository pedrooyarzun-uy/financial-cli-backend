package dto

type CreateReq struct {
	Name     string  `json:"name" binding:"required"`
	Number   string  `json:"number" binding:"required"`
	Currency int     `json:"currency" binding:"required"`
	Cash     float64 `json:"cash" binding:"cash"`
	Bank     int     `json:"bank" binding:"bank"`
	Owner    int     `json:"owner" binding:"owner"`
}

type DeleteReq struct {
	Id int `json:"id" binding:"required"`
}
