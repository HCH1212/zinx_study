package ziface

// 包装客户端请求的链接信息和请求数据
type IRequest interface {
	GetConnection() IConnection // 得到当前链接
	GetData() []byte            // 得到请求的消息数据
}
