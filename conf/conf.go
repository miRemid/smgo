package conf

import (
	"github.com/joho/godotenv"
)

// Init 环境初始化
func Init() {
	// 读取本地环境
	godotenv.Load()
}