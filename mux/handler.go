package mux

import "net/http"

// 对如何处理逻辑和存储路由的能力抽象
type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	Route(method, pattern string, handlerFunc http.HandlerFunc)
}
