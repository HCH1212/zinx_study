package ziface

import "net"

// 链接模块，处理conn
type IConnection interface {
	Start()                         // 启动链接 让当前链接准备开始工作
	Stop()                          // 停止链接 结束当前链接的工作
	GetTCPConnection() *net.TCPConn // 获取当前链接的socket
	GetConnID() uint32              // 获取当前链接id
	RemoteAddr() net.Addr           // 获取远程客户端地址
	Send(data []byte) error         // 发送数据到远程客户端
}

// 处理链接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
