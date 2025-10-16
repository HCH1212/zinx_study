package ziface

// 路由接口（可以根据消息来对应不同的路由），路由的数据都是IRequest
type IRouter interface {
	PreHandle(request IRequest)  // 在处理业务之前的钩子函数hook
	Handle(request IRequest)     // 处理业务的主方法
	PostHandle(request IRequest) // 在处理业务之后的钩子函数hook
}
