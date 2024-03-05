package main

import (
	"fmt"
	"net/http"
)

// HTTP服务端配置
// 业务请求相应处理
func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Println("hello")
	fmt.Fprintln(res, "<h1>welcome</h1>")
}
func main() {
	//路由到指定站点位置跳转相关函数处理
	fmt.Println("服务端启动了。。。。。")
	http.HandleFunc("/", hello)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println("启动监听失败,err:", err)
		return
	}
}
