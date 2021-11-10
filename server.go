package mux

import (
	"fmt"
	"net/http"
)

// 对 Server 行为的抽象
type Server interface {
	Group(prefix string) *RouterGroup
	Start(address string) error
}

// 对 Server 的一种实现
// handler 处理 ServeHTTP Route
type miniHTTPServer struct {
	Name    string
	handler Handler
	// http 请求处理链
	root HandlerFunc
	*RouterGroup
	groups []*RouterGroup
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
	server := &miniHTTPServer{
		Name:    name,
		handler: handler,
		root:    root,
	}
	server.RouterGroup = &RouterGroup{server: server}
	return server
}
