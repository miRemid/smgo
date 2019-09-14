package models

// Image 上传序列化结构
type Image struct {
	BaseResponse
	Data		ImageData	`json:"data"`	
}

// Images 多数据结构
type Images struct {
	BaseResponse
	Data		[]ImageData	`json:"data"`
}

// ImageData 上传数据信息
type ImageData struct {
	Delete		string	`json:"delete"`
	FileID		int		`json:"file_id"`
	FileName	string	`json:"filename"`
	Hash		string	`json:"hash"`
	Height		int		`json:"height"`
	Page 		string	`json:"page"`
	Path		string	`json:"path"`
	Size		int		`json:"size"`
	StoreName	string	`json:"storename"`
	URL			string	`json:"url"`
	Width		int		`json:"width"`
}
