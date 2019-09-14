package models

// Profile 用户序列化结构
type Profile struct {
	BaseResponse
	Data	ProfileData	`json:"data"`
}

// ProfileData 用户信息
type ProfileData struct {
	Username	string	`json:"username"`	
	Role		string	`json:"role"`
	GroupExpire	string	`json:"group_expire"`
	DiskUsage	string	`json:"disk_usage"`
	DiskLimit	string	`json:"disk_limit"`
}