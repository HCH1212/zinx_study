package znet

import (
	"fmt"
	"net"
	"zinx_study/ziface"
)

type Connection struct {
	Conn      *net.TCPConn      // 当前链接的socket
	ConnID    uint32            // 当前链接id
	isClosed  bool              // 当前链接状态
	handleAPI ziface.HandleFunc // 当前链接所绑定的处理业务方法
	ExitChan  chan struct{}     // 告知当前链接已经停止的channel
}

func NewConnection(conn *net.TCPConn, connID uint32, callBackAPI ziface.HandleFunc) ziface.IConnection {
	return &Connection{
		Conn:      conn,
		ConnID:    connID,
		handleAPI: callBackAPI,
		isClosed:  false,
		ExitChan:  make(chan struct{}),
	}
}

func (c *Connection) Start() {
	fmt.Println("starting... ", c.ConnID)

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
