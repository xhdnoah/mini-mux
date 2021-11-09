package mux

import (
	"net/http"
	"sync"
)

var _ Handler = &HandlerBasedOnMap{}

type HandlerBasedOnMap struct {
	handlers sync.Map
	BaseHandler
}

func (h *HandlerBasedOnMap) ServeHTTP(c *Context) {
	r, w := c.Request, c.Writer
	key := h.key(r.Method, r.URL.Path)
	handler, ok := h.handlers.Load(key)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("no handler registered"))
		return
	}

	handler.(HandlerFunc)(c)
}

func (h *HandlerBasedOnMap) Route(method, pattern string, handlerFunc HandlerFunc) {
	key := h.key(method, pattern)
	h.handlers.Store(key, handlerFunc)
}

func NewHandlerBasedOnMap() *HandlerBasedOnMap {
	return &HandlerBasedOnMap{}
}
