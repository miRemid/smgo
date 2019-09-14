package smgo

import (
	"net/http"
	"encoding/json"

	"github.com/miRemid/smgo/models"
)

// Clear 清除上传历史
func (sm *SmClient) Clear() (models.BaseResponse, error) {
	var response models.BaseResponse
	// 2. 构造请求
	req, err := http.NewRequest("GET", ClearURL, nil)
	if err != nil {		
		return response, err
	}
	// 添加header
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// 3. 发送请求
	res, err := sm.HTTPClient.Do(req)
	if err != nil{
		return response, err
	}
	defer res.Body.Close()
	// 4. 格式化结果
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return response, err
	}
	return response, nil
}