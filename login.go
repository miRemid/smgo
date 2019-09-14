package smgo

import (
	"fmt"
	"strings"	
	"net/http"
	"encoding/json"	

	"github.com/miRemid/smgo/models"
)

// Login 获得用户token
func (sm *SmTokenClient) Login() (models.LoginRes, error) {
	var token models.LoginRes
	// 1. 构造参数
	tmp := fmt.Sprintf("username=%s&password=%s", sm.UserName, sm.Password)	
	params := strings.NewReader(tmp)
	// 2. 构造请求
	req, err := http.NewRequest("POST", TokenURL, params)
	if err != nil {		
		return token, err
	}
	// 添加header
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// 3. 发送请求
	res, err := sm.HTTPClient.Do(req)
	if err != nil {		
		return token, err
	}
	// 关闭body
	defer res.Body.Close()
	// 4. 解析body到结构体
	err = json.NewDecoder(res.Body).Decode(&token)
	if err != nil {		
		return token, err
	}
	// 设置登陆状态
	sm.login = true
	sm.Token = token.Data.Token	
	return token, nil
}