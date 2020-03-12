# Smgo!
Smgo!是一个基于Golang的SM.MS图床第三方Cli&SDK工具

# 安装
```sh
go get github.com/miRemid/smgo.git
export PATH=$PATH:`你的GOPATH`/bin
```
# Cli 使用
## 上传
```shell
$ smgo u -h
NAME:
   smgo upload - upload image to sm.ms

USAGE:
   smgo upload [command options] [arguments...]

OPTIONS:
   --token value, -t value   set account token
   --time value, --tm value  set timeout (default: 5)
   --help, -h                show help (default: false)
$ smgo u abc.jpg 
{
  "success": true,
  "code": "success",
  "message": "Upload success.",
  "RequestId": "FE9A9889-FD9E-47FD-8E48-D8CF55B8883A",
  "data": {
    "delete": "https://sm.ms/delete/bPSMvGAk8rUK96ZjsfzYJiEa2T",
    "file_id": 0,
    "filename": "abc.jpg",
    "hash": "bPSMvGAk8rUK96ZjsfzYJiEa2T",
    "height": 1080,
    "page": "https://sm.ms/image/Y2jn4T5PvFHxQlN",
    "path": "/2020/03/12/Y2jn4T5PvFHxQlN.jpg",
    "size": 336198,
    "storename": "Y2jn4T5PvFHxQlN.jpg",
    "url": "https://i.loli.net/2020/03/12/Y2jn4T5PvFHxQlN.jpg",
    "width": 1920
 }
}
```
## 删除
```shell
$ smgo d -h                        
NAME:
   smgo delete - delete image from sm.ms

USAGE:
   smgo delete [command options] [arguments...]

OPTIONS:
   --token value, -t value   set account token
   --time value, --tm value  set timeout (default: 5)
   --help, -h                show help (default: false)
$ smgo d bPSMvGAk8rUK96ZjsfzYJiEa2T
{
  "success": true,
  "code": "success",
  "message": "File delete success.",
  "data": [],
  "RequestId": "2589D55B-B432-4367-BA63-920435AA4D94"
}
```
## 查看个人信息
```sh
$ smgo p -h
NAME:
   smgo profile - print profile infomation

USAGE:
   smgo profile [command options] [arguments...]

OPTIONS:
   --token value, -t value   set account token
   --time value, --tm value  set timeout (default: 5)
   --help, -h                show help (default: false)
$ smgo p -t 你的账号Token
{
  "success": true,
  "code": "success",
  "message": "Get user profile success.",
  "RequestId": "16CC6A16-27FD-49E4-BCEB-A25740F837C0",
  "data": {
    "username": "hahaha",
    "role": "user",
    "group_expire": "0000-00-00",
    "disk_usage": "4.09 MB",
    "disk_limit": "5.00 GB"
  }
}
```
## 查看上传记录
```sh
$ smgo h -h                                 
NAME:
   smgo history - print upload history

USAGE:
   smgo history [command options] [arguments...]

OPTIONS:
   --token value, -t value   set account token
   --time value, --tm value  set timeout (default: 5)
   --help, -h                show help (default: false)
```
## 清除上传记录
```sh
$ smgo c -h                                 
NAME:
   smgo clear - print upload clear

USAGE:
   smgo clear [command options] [arguments...]

OPTIONS:
   --token value, -t value   set account token
   --time value, --tm value  set timeout (default: 5)
   --help, -h                show help (default: false)
```

# SDK Usage
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