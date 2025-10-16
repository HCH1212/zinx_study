package znet

import "zinx_study/ziface"

// 默认router
type BaseRouter struct{}

// 方法都留空，待使用时根据需要实现
func (br *BaseRouter) PreHandle(request ziface.IRequest) {}

func (br *BaseRouter) Handle(request ziface.IRequest) {}

func (br *BaseRouter) PostHandle(request ziface.IRequest) {}
