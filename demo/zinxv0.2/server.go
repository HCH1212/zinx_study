package main

import "zinx_study/znet"

// 基于zinx来开发的服务器应用程序 测试
func main() {
	// 1. 获取句柄
	s := znet.NewServer("[zinx v0.1]")
	// 2. 启动
	s.Serve()
}
