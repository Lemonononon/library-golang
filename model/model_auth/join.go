package model_auth

type JoinReq struct {
	Password  string `json:"password" binding:"required"`
	AdminName string `json:"admin_name" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
}
type JoinResp struct {
	Token string `json:"token"`
}
