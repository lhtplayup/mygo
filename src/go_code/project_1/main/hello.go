package main

import (
	"fmt"
	"net"
	"net/http"
)

// 获取本机网卡IP
func getLocalIP() (ipv4 string, err error) {
	// 获取所有网卡
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	// 取第一个非lo的网卡IP
	for _, addr := range addrs {
		// 这个网络地址是IP地址: ipv4, ipv6
		ipNet, isIpNet := addr.(*net.IPNet)
		if isIpNet && !ipNet.IP.IsLoopback() {
			// 跳过IPV6
			if ipNet.IP.To4() != nil {
				ipv4 = ipNet.IP.String() // 192.168.1.1
				return
			}
		}
	}

	return
}

// HTTP服务端配置
// 业务请求相应处理
func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Println("hello")
	fmt.Fprintln(res, "<h1>welcome</h1>")
}
func main() {
	//路由到指定站点位置跳转相关函数处理
	ipv4, _ := getLocalIP()
	fmt.Printf("ipv4 is <%s>", ipv4)
	fmt.Println("服务端启动了。。。。。")
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(ipv4+":8080", nil)
	if err != nil {
		fmt.Println("启动监听失败,err:", err)
		return
	}
}
