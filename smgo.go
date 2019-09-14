package smgo

import (
	"net/http"
	"errors"
	"time"
	"fmt"
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
}

// SmTokenClient 主程序
type SmTokenClient struct {
	SmClient
	UserName	string
	Password	string
	Token		string
	login		bool
}

// NewSmClient 创建一个新的SmClient
func NewSmClient() SmClient {
	return SmClient{}
}

// NewSmClientWithToken 创建一个登陆Client
func NewSmClientWithToken(username, password string) SmTokenClient {
	sm := SmTokenClient{
		UserName: 	username,
		Password: 	password,
	}
	sm.HTTPClient.Timeout = time.Second * 10
	if token, err := sm.Login(); err != nil{
		fmt.Println(err)
	}else{
		fmt.Printf("Login Success, Token = %s\n", token.Data.Token)
	}
	return sm
}

// CheckLogin 检查登陆
func (sm *SmTokenClient) CheckLogin() error {
	if !sm.login {
		return	errors.New("Need Login")
	}
	return nil
}