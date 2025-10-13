package dto

type SignInReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInRes struct {
	Message string `json:"message"`
	Auth    string `json:"auth"`
}

type SignUpReq struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpRes struct {
	Message string `json:"message"`
}
