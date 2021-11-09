package mux

import "fmt"

type BaseHandler struct{}

func (h *BaseHandler) key(method, path string) string {
	return fmt.Sprintf("%s-%s", method, path)
}
