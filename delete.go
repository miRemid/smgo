package smgo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"smgo/models"
)

// Delete 删除指定文件
func (sm *SmClient) Delete(hash string) (models.BaseResponse, error) {
	var response models.BaseResponse
	// 1. 构造url
	url := DeleteURL + hash
	// 2. 构造请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return response, err
	}
	// 添加header
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// 3. 发送请求
	res, err := sm.HTTPClient.Do(req)
	if err != nil {
		return response, err
	}
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	// 4. 格式化结果
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
