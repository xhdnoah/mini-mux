package main

import (
	"fmt"
	"mux"
	"net/http"
)

func hello(c *mux.Context) {
	fmt.Println("Hello World!")
}

func login(c *mux.Context) {
	c.JSON(http.StatusOK, mux.H{
		"username": c.PostForm("username"),
		"password": c.PostForm("password"),
	})
}

func main() {
	s := mux.NewMiniHTTPServer("alice")
	s.Route("GET", "/hello", hello)
	s.Route("POST", "/login", login)
	s.Start(":8000")
}
