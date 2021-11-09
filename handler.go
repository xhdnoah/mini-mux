package mux

// 对如何处理逻辑和存储路由的能力抽象
type Handler interface {
	ServeHTTP(c *Context)
	Route(method, pattern string, handlerFunc HandlerFunc)
}

type HandlerFunc func(*Context)
