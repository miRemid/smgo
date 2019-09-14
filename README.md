# Smgo!

Smgo!是一个基于Golang的SM.MS图床第三方SDK工具，方便土拨鼠们在写博客的时候能够快速得到图片的图床连接，加快文档编写速度.

# Usage
### install安装
```
git clone https://github.com/miRemid/smgo.git 到你的GOPATH中
```
### First Try初步使用
```golang
package main
import "fmt"
import "github.com/miRemid/smgo"

func main(){
    // 1. 可以选择创建非用户客户端或用户客户端
    // 不带token客户端
    // sm := smgo.NewClient()
    // 推荐注册账号使用
    sm := smgo.NewTokenClient(username, password)

    // 上传文件
    if res, err := sm.Upload(filepath); err != nil {
        Handle(err)
    }else{
        fmt.Println(res.Data.URL)
    }
}
```