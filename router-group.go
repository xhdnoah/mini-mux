package mux

import "log"

type Routable interface {
	Route(method string, path string, handlerFunc HandlerFunc)
}

type RouterGroup struct {
	Routable
	prefix string
	parent *RouterGroup
	server *miniHTTPServer
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	server := group.server
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		server: server,
	}
	server.groups = append(server.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) Route(method, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.server.handler.Route(method, pattern, handler)
}
