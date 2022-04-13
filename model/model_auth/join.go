package model_auth

type JoinReq struct {
	AdminName string `json:"admin_name" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
}
type JoinResp struct {
	Token string `json:"token"`
}
