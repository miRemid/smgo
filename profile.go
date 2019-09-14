package smgo

import (	
	"net/http"
	"encoding/json"

	"github.com/miRemid/smgo/models"
)

// Profile 获取用户个人信息
func (sm *SmTokenClient) Profile() (models.Profile, error) {
	var profile models.Profile
	// 检查登陆状态
	if err := sm.CheckLogin(); err != nil{
		return profile, err
	}	
	// 1. 构造请求
	req, err := http.NewRequest("POST", ProfileURL, nil)
	if err != nil {
		return profile, err
	}
	// 添加header
	req.Header.Add("content-type", "multipart/form-data")
	req.Header.Add("Authorization", sm.Token)
	// 2. 发送请求
	res, err := sm.HTTPClient.Do(req)
	if err != nil {
		return profile, err
	}
	defer res.Body.Close()
	// 3. 过滤结果
	err = json.NewDecoder(res.Body).Decode(&profile)	
	if err != nil {
		return profile, err
	}	
	return profile, nil
}