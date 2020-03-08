package form

type LoginForm struct {
	Username    string `json:"username" form:"username" binding:"required"`
	Password    string `json:"password" form:"password" binding:"required"`
	VerifyCode  string `json:"verifyCode" form:"verifyCode"`
	AccountType string `json:"accountType" form:"accountType"`
}
