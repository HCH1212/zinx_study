package znet

import (
	"fmt"
	"net"
	"zinx_study/ziface"
)

// IServer的接口实现
type Server struct {
	Name      string // 服务器名称
	IPVersion string // 服务器绑定的ip版本（v4 or v6）
	IP        string // 服务器监听的ip
	Port      int    // 服务器监听的端口
}

// 初始化模块
func NewServer(name string) ziface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
}

// 启动服务器
func (s *Server) Start() {
	fmt.Printf("[Start] IP:%s, Port:%d\n", s.IP, s.Port)

	go func() {

		// 1， 获取一个tcp的addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println(err)
			return
		}

		// 2. 监听服务器得地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println(err)
			return
		}

		// 3. 阻塞的等待客户端连接，处理客户端连接业务（读写）
		for {
			// 有客户端连接
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println(err)
				continue
			}

			// 以及与客户端建立连接，做业务（回显）
			go func() {
				for {
					buf := make([]byte, 512)
					n, err := conn.Read(buf)
					if err != nil {
						fmt.Println(err)
						continue
					}

					fmt.Println("get: ", string(buf[:n]))

					if _, err = conn.Write(buf[:n]); err != nil {
						fmt.Println(err)
						continue
					}
				}
			}()
		}

	}()
}

// 停止服务器
func (s *Server) Stop() {
	// TODO 将一些服务器的资源、状态或连接信息 进行停止或回收
}

// 运行服务器
func (s *Server) Serve() {
	// 启动server的服务功能
	s.Start()

	// TODO 这里可以做一些额外业务

	// 阻塞状态
	select {}
}
