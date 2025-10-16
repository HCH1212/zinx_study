package znet

import (
	"fmt"
	"net"
	"zinx_study/ziface"
)

type Connection struct {
	Conn     *net.TCPConn   // 当前链接的socket
	ConnID   uint32         // 当前链接id
	isClosed bool           // 当前链接状态
	ExitChan chan struct{}  // 告知当前链接已经停止的channel
	Router   ziface.IRouter // 当前链接的router业务方法
}

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) ziface.IConnection {
	return &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		ExitChan: make(chan struct{}),
		Router:   router,
	}
}

func (c *Connection) StartReader() {
	fmt.Println("reader goroutine is running...")

	defer c.Stop()
	for {
		buf := make([]byte, 512)
		n, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// 获取request
		req := Request{
			conn: c,
			data: buf[:n],
		}
		// 获取数据后，调用对应路由方法
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}

func (c *Connection) Start() {
	fmt.Println("starting... ", c.ConnID)

	// 启动从当前链接的读数据业务
	go c.StartReader()
	// TODO 启动从当前链接的写数据业务
}

func (c *Connection) Stop() {
	fmt.Println("stopping...")

	if c.isClosed {
		return
	}
	c.isClosed = true
	c.Conn.Close()
	close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	_, err := c.Conn.Write(data)
	return err
}
