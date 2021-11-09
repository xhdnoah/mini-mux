package mux

import (
	"net/http"
	"strings"
	"sync"
)

type HandlerBasedOnTree struct {
	roots    sync.Map
	handlers sync.Map
	BaseHandler
}

var supportMethods = [4]string{
	http.MethodGet,
	http.MethodPost,
	http.MethodPut,
	http.MethodDelete,
}

// end with the first *
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, v := range vs {
		if v != "" {
			parts = append(parts, v)
			if v[0] == '*' {
				break
			}
		}
	}

	return parts
}

func (h *HandlerBasedOnTree) Route(method, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	key := h.key(method, pattern)
	rootNode, ok := h.roots.Load(method)
	if !ok {
		rootNode = &node{}
		h.roots.Store(method, rootNode)
	}
	rootNode.(*node).insert(pattern, parts, 0)
	h.handlers.Store(key, handler)
}

// path: /assets/js/vendor.js, pattern: /assets/*filepath
func (h *HandlerBasedOnTree) getRouteNode(method, path string) (*node, map[string]string) {
	pathParts := parsePattern(path)
	// * : 路由参数
	params := make(map[string]string)
	root, ok := h.roots.Load(method)

	if !ok {
		return nil, nil
	}

	n := root.(*node).search(pathParts, 0)
	if n != nil {
		patternParts := parsePattern(n.pattern)
		for index, part := range patternParts {
			if part[0] == ':' {
				params[part[1:]] = pathParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(pathParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (h *HandlerBasedOnTree) getRouteNodes(method string) []*node {
	root, ok := h.roots.Load(method)
	if !ok {
		return nil
	}
	nodes := make([]*node, 0)
	root.(*node).travel(&nodes)
	return nodes
}

func (h *HandlerBasedOnTree) ServeHTTP(c *Context) {
	method, path := c.Request.Method, c.Request.URL.Path
	n, params := h.getRouteNode(method, path)
	if n != nil {
		c.Params = params
		key := h.key(method, path)
		f, _ := h.handlers.Load(key)
		f.(HandlerFunc)(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", path)
	}
}

func NewHandlerBasedOnTree() *HandlerBasedOnTree {
	return &HandlerBasedOnTree{}
}
