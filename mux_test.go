package mux

import (
	"fmt"
	"reflect"
	"testing"
)

func testRoute(r Routable) {
	r.Route("GET", "/", nil)
	r.Route("GET", "/hello/:name", nil)
	r.Route("GET", "/a/b/c", nil)
	r.Route("GET", "/assets/*filepath", nil)
}

func newTestTreeRouter() *HandlerBasedOnTree {
	t := NewHandlerBasedOnTree()
	testRoute(t)
	return t
}

func newTestGroupRouter() []*RouterGroup {
	s := NewMiniHTTPServer("alice")
	v1 := s.Group("/v1")
	testRoute(v1)
	v2 := v1.Group("/v2")
	{
		v2.Route("POST", "/login", nil)
	}
	return []*RouterGroup{v1, v2}
}

func TestNestedGroup(t *testing.T) {
	vs := newTestGroupRouter()
	if vs[1].prefix != "/v1/v2" {
		t.Fatal("v2 prefix should be /v1/v2")
	}
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
