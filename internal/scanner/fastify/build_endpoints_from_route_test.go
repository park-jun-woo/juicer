//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what buildEndpointsFromRoute 테스트
package fastify

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
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

func TestBuildEndpointsFromRoute_All(t *testing.T) {
	r := routeInfo{Method: "all", Path: "/", Handler: "h"}
	eps := buildEndpointsFromRoute(r, "", "routes/x.ts", []byte(""), map[string]*sitter.Node{})
	if len(eps) != 5 {
		t.Fatalf("expected 5 endpoints for 'all', got %d", len(eps))
	}
}
