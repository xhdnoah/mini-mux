package mux

import (
	"fmt"
	"net/http"
	"sync"
)

var _ Handler = &HandlerBasedOnMap{}

type HandlerBasedOnMap struct {
	handlers sync.Map
}

func (h *HandlerBasedOnMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := h.key(r.Method, r.URL.Path)
	handler, ok := h.handlers.Load(key)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("no handler registered"))
		return
	}

	handler.(http.HandlerFunc)(w, r)
}

func (h *HandlerBasedOnMap) Route(method string, pattern string, handlerFunc http.HandlerFunc) {
	key := h.key(method, pattern)
	h.handlers.Store(key, handlerFunc)
}

func (h *HandlerBasedOnMap) key(method string, path string) string {
	return fmt.Sprintf("%s-%s", method, path)
}

func NewHandlerBasedOnMap() *HandlerBasedOnMap {
	return &HandlerBasedOnMap{}
}
