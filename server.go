package mux

import (
	"fmt"
	"net/http"
)

// 对 Server 行为的抽象
type Server interface {
	Start(address string) error
	Route(method, path string, handler HandlerFunc)
}

// 对 Server 的一种实现
// handler 处理 ServeHTTP Route
type miniHTTPServer struct {
	Name    string
	handler Handler
	// http 请求处理链
	root HandlerFunc
}

func (s *miniHTTPServer) Route(method, pattern string, handlerFunc HandlerFunc) {
	// 代理到 Route 实现
	s.handler.Route(method, pattern, handlerFunc)
}

func (s *miniHTTPServer) Start(address string) error {
	fmt.Println(s.Name + " is listening on " + address[1:])
	return http.ListenAndServe(address, s)
}

func (s *miniHTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContext(w, r)
	s.root(c)
}

func NewMiniHTTPServer(name string) Server {
	handler := NewHandlerBasedOnTree()
	root := handler.ServeHTTP
	return &miniHTTPServer{
		Name:    name,
		handler: handler,
		root:    root,
	}
}
