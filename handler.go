package mux

// 对如何处理逻辑和存储路由的能力抽象
type Handler interface {
	Routable
	ServeHTTP(c *Context)
}

type HandlerFunc func(*Context)
