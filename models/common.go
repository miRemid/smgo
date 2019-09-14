package models

import "encoding/json"

// BaseResponse SmApi返回信息结构
type BaseResponse struct {
	Success		bool		`json:"success"`
	Code		string		`json:"code"`
	Message		string		`json:"message"`
	Data		interface{}	`json:"data"`
	RequestID	string		`json:"RequestId"`
}

func (b *BaseResponse) String() string {
	data, _ := json.Marshal(b)
	return string(data)
}