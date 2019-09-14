package smgo

import (
	"net/http"
	"encoding/json"	

	"github.com/miRemid/smgo/models"
)

// History 获取上传历史
func (sm *SmClient) History() (models.Images, error) {
	var imgs models.Images
	// 1. 构造请求
	req, err := http.NewRequest("GET", HistoryURL, nil)
	if err != nil {
		return imgs, err
	}
	// 2. 发送请求
	res, err := sm.HTTPClient.Do(req)
	if err != nil {
		return imgs, err
	}
	defer res.Body.Close()
	// 解析
	err = json.NewDecoder(res.Body).Decode(&imgs)	
	if err != nil {
		return imgs, err
	}
	return imgs, nil
}

// History 获取用户上传历史
func (sm *SmTokenClient) History() (models.Images, error) {
	var imgs models.Images
	// 1. 构造请求
	req, err := http.NewRequest("POST", UhistoryURL, nil)
	if err != nil {
		return imgs, err
	}
	// 检查登陆状态
	if err = sm.CheckLogin(); err != nil {
		return imgs, err
	}
	req.Header.Add("Authorization", sm.Token)
	// 2. 发送请求
	res, err := sm.HTTPClient.Do(req)
	if err != nil {
		return imgs, err
	}
	defer res.Body.Close()
	// 解析	
	err = json.NewDecoder(res.Body).Decode(&imgs)	
	if err != nil {
		return imgs, err
	}
	return imgs, nil
}