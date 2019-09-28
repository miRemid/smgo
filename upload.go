package smgo

import (
	"io"
	"os"
	"bytes"
	"net/http"
	"encoding/json"
	"mime/multipart"

	"github.com/miRemid/smgo/models"
)

// newFileUploadRequest 生成请求
func (sm *SmClient) newFileRequest(url, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("smfile", path)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	if err = writer.Close(); err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	if sm.CheckLogin(){
		req.Header.Add("Authorization", sm.Token)
	}
	return req, err
}

// Upload 上传图片
func (sm *SmClient) Upload(filePath string) (models.Image, error) {
	var img models.Image
	// 1. 构造请求
	req, err := sm.newFileRequest(UploadURL, filePath)
	if err != nil{
		return img, nil
	}
	// 3. 发送请求
	res, err := sm.HTTPClient.Do(req)
	if err != nil {
		return img, nil
	}
	defer res.Body.Close()
	// 4. 解析数据
	err = json.NewDecoder(res.Body).Decode(&img)	
	if err != nil {
		return img, err
	}	
	return img, nil
}

// Uploads 批量上传
func (sm *SmClient) Uploads(filePath ...string) ([]models.Image, error) {
	var imgs []models.Image
	for _, path := range filePath{
		if img, err := sm.Upload(path); err != nil {
			return imgs, err
		}else{
			imgs = append(imgs, img)
		}
	}
	return imgs, nil
}
