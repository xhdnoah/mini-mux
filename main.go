package main

import (
	"fmt"
	"mux"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World!")
}

func main() {
	s := mux.NewMiniHTTPServer("alice")
	s.Route("GET", "/hello", hello)
	s.Start(":8000")
}
