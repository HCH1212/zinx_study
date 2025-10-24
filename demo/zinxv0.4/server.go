package main

import (
	"fmt"
	"zinx_study/ziface"
	"zinx_study/znet"
)

// 基于zinx来开发的服务器应用程序 测试
func main() {
	// 1. 获取句柄
	s := znet.NewServer("[zinx v0.4]")
	// 2. 添加路由
	s.AddRouter(&PingRouter{})
	// 3. 启动
	s.Serve()
}

// ping 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

func (pr *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("PreHandle...")

	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (pr *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Handle...")

	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping ping ping...\n"))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (pr *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("PostHandle...")

	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println(err)
		return
	}
}
