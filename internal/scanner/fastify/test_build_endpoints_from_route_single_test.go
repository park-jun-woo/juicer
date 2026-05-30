//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestBuildEndpointsFromRoute_Single 테스트
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestBuildEndpointsFromRoute_Single(t *testing.T) {
	r := routeInfo{Method: "GET", Path: "/:id", Handler: "getUser"}
	eps := buildEndpointsFromRoute(r, "/users", "routes/users.ts", []byte(""), map[string]*sitter.Node{})
	if len(eps) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(eps))
	}
	ep := eps[0]
	if ep.Method != "GET" {
		t.Errorf("method = %q", ep.Method)
	}
	if ep.Path != "/users/{id}" {
		t.Errorf("path = %q, want /users/{id}", ep.Path)
	}
	if ep.Handler != "getUser" {
		t.Errorf("handler = %q", ep.Handler)
	}
}
