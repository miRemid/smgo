package models

// LoginRes 登陆结构
type LoginRes struct {
	BaseResponse
	Data	LoginData	`json:"data"`
}

// LoginData 登陆结果信息
type LoginData struct {
	Token		string	`json:"token"`
}
