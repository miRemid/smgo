package smgo

import (
	"net/http"
	"time"
)

const (
	// BaseURL 基本域名
	BaseURL			=	"https://sm.ms"
	// TokenURL 登陆						
	TokenURL 		= 	BaseURL + "/api/v2/token"		
	// ProfileURL 个人信息
	ProfileURL 		=	BaseURL + "/api/v2/profile"
	// DeleteURL 删除图片
	DeleteURL 		=	BaseURL + "/api/v2/delete/"
	// ClearURL 清楚历史
	ClearURL		=	BaseURL + "/api/v2/clear"		
	// UhistoryURL 获取历史(非登陆)
	UhistoryURL		=	BaseURL + "/api/v2/upload_history"	
	// HistoryURL 获取历史(登陆)
	HistoryURL		=	BaseURL + "/api/v2/history"
	// UploadURL 上传文件
	UploadURL		=	BaseURL + "/api/v2/upload"
)

// SmClient 默认客户端
type SmClient struct {
	HTTPClient 	http.Client
	Token	string
}

// NewSmClient 创建一个新的SmClient
func NewSmClient() *SmClient {
	return &SmClient{}
}

// SetTimeout 设置超时
func (sm *SmClient) SetTimeout(timeout time.Duration){
	sm.HTTPClient.Timeout = timeout
}

// SetToken 设置token
func (sm *SmClient) SetToken(token string) {
	sm.Token = token
}

// CheckLogin 检查登陆
func (sm *SmClient) CheckLogin() bool {	
	res := sm.Token != ""
	return res
}
