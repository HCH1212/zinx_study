package main

import (
	"fmt"
	"net"
	"time"
)

// 模拟客户端
func main() {
	fmt.Println("client start...")
	time.Sleep(1 * time.Second)

	// 1. 直接连接远程服务器，获取conn
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		// 2. 调用write写数据
		_, err = conn.Write([]byte("hello zinx v0.2"))
		if err != nil {
			fmt.Println(err)
			return
		}

		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("echo: ", string(buf[:n]))

		// cpu阻塞
		time.Sleep(1 * time.Second)
	}
}
