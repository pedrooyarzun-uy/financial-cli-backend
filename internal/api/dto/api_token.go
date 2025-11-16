package dto

import (
	"time"
)

type CreateApiTokenReq struct {
	Keyword string `json:"keyword" binding:"required"`
	Name    string `json:"name" binding:"required"`
}

type ApiTokenResponse struct {
	Id        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	Revoked   bool      `db:"revoked"`
}

type GetAllApiTokenRes struct {
	Message string             `json:"message"`
	Tokens  []ApiTokenResponse `json:"tokens"`
}
