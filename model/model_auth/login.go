package model_auth

type LoginReq struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResp struct {
	Token string `json:"token"`
}
