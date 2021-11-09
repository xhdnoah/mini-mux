package mux

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestTreeRouter() *HandlerBasedOnTree {
	t := NewHandlerBasedOnTree()
	t.Route("GET", "/", nil)
	t.Route("GET", "/hello/:name", nil)
	t.Route("GET", "/a/b/c", nil)
	t.Route("GET", "/assets/*filepath", nil)
	return t
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRouteParam(t *testing.T) {
	r := newTestTreeRouter()
	n, ps := r.getRouteNode("GET", "/hello/world")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	if ps["name"] != "world" {
		t.Fatal("nae should be equal to 'world'")
	}

	fmt.Printf("matched path: %s, params['name']: %s\n", n.pattern, ps["name"])
}

func TestGetRoutePath(t *testing.T) {
	r := newTestTreeRouter()
	n, ps := r.getRouteNode("GET", "/assets/html/test.html")
	ok := n.pattern == "/assets/*filepath" && ps["filepath"] == "html/test.html"
	if !ok {
		t.Fatal("pattern shoule be /assets/*filepath & filepath shoule be html/test.html")
	}
}

func TestGetRoutes(t *testing.T) {
	r := newTestTreeRouter()
	nodes := r.getRouteNodes("GET")
	for i, n := range nodes {
		fmt.Println(i+1, n)
	}

	if len(nodes) != 4 {
		t.Fatal("the number of routes should be 4")
	}
}
