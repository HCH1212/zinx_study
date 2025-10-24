package znet

import (
	"fmt"
	"net"
	"zinx_study/utils"
	"zinx_study/ziface"
)

// IServer的接口实现
type Server struct {
	Name      string         // 服务器名称
	IPVersion string         // 服务器绑定的ip版本（v4 or v6）
	IP        string         // 服务器监听的ip
	Port      int            // 服务器监听的端口
	Router    ziface.IRouter // 当前server的router
}

// 初始化模块
func NewServer(name string) ziface.IServer {
	return &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		Router:    nil,
	}
}

// 启动服务器
func (s *Server) Start() {
	fmt.Printf("[Zinx] 配置信息 %+v\n", utils.GlobalObject)

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

		var cid uint32 = 0 // 链接id

		// 3. 阻塞的等待客户端连接，处理客户端连接业务（读写）
		for {
			// 有客户端连接
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println(err)
				continue
			}

			// 有新链接就进行绑定
			dealConn := NewConnection(conn, cid, s.Router)
			cid++

			go dealConn.Start()
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

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("add router success!")
}
