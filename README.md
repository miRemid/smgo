# Smgo!

Smgo!是一个基于Golang的SM.MS图床第三方SDK工具

# Install
```
git clone https://github.com/miRemid/smgo.git
```
# Usage
```golang
package main
import "fmt"
import "github.com/miRemid/smgo"

// 请注意请求不要过于频繁
func main(){
    sm := smgo.NewSmClient()
    // 可选是否设置token，不设置token则为匿名上传
    // 1. 可以选择登陆设置token，登陆会返回用户token
    token, _ := sm.Login("username", "password")
    // 2. 也可以选择直接设置token
    sm.SetToken("token")

    // 查看上传历史
	if res, err := sm.History(); err != nil {
		fmt.Println(err)
	}else {
		for _, v := range res.Data{
			fmt.Printf("Name:%s, URL:%s\n", v.FileName, v.URL)
		}
    }
    
    // 上传一个文件
    res, err := sm.Upload("76336523.jpg")
	if err != nil {
		log.Println(err)
	}else{
		log.Printf("%v", res)
    }
    
    // 上传文件流
    file, err := os.Open("76336523.jpg")
    defer file.Close()	
	if err != nil {
		log.Println(err)
	}
	res, err := sm.UploadStream(file, "filename")
    if err != nil {
		log.Println(err)
	}else{
		log.Printf("%v", res)
    }

    // 上传多个文件
    files := []string{"1.jpg", "2.jpg"}
    res, err := sm.Uploads(files...)
	if err != nil {
		log.Println(err)
	}else{
		log.Printf("%v", res)
    }

    // 删除文件，参数为文件hash值，在上传文件的Data中可以找到
    sm.Delete("文件的hash值")
}
```