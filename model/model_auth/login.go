package model_auth

type LoginReq struct {
	AdminID  string `json:"admin_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginResp struct {
	Token string `json:"token"`
}