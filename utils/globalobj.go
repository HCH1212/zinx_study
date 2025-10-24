package utils

import (
	"encoding/json"
	"io/ioutil"
	"zinx_study/ziface"
)

// 配置
type GlobalObj struct {
	TcpServer ziface.IServer // 当前zinx全局的Server对象
	Host      string         // 当前服务器主机监听的ip
	TcpPort   int            // 当前服务器主机监听的端口
	Name      string         // 当前服务器名字

	Version        string // 当前zinx版本号
	MaxConn        int    // 当前服务器允许的最大连接数
	MaxPackageSize uint32 // 当前zinx数据包的最大值
}

// 提供全局可访问配置
var GlobalObject *GlobalObj

// 初始化GlobalObject
func init() {
	// 不加载配置文件的默认值
	GlobalObject = &GlobalObj{
		Name:           "ZinxServerApp",
		Version:        "V0.4",
		TcpPort:        8999,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	// 从conf/zinx.json加载配置
	GlobalObject.Reload()
}

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("demo/zinxv0.4/conf/zinx.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, GlobalObject)
	if err != nil {
		panic(err)
	}
}
