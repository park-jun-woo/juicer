//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what buildOneEndpoint 테스트
package fastify

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestBuildOneEndpoint_WithRequest(t *testing.T) {
	r := routeInfo{Method: "GET", Handler: "h", Line: 7}
	ep := buildOneEndpoint("GET", "/users/{id}", r, "routes/u.ts", []string{"id"}, []byte(""), map[string]*sitter.Node{})
	if ep.Method != "GET" || ep.Path != "/users/{id}" || ep.Handler != "h" || ep.Line != 7 {
		t.Fatalf("endpoint fields wrong: %+v", ep)
	}
	if ep.Request == nil || len(ep.Request.PathParams) != 1 {
		t.Fatalf("expected path params in request, got %v", ep.Request)
	}
}

func TestBuildOneEndpoint_NoRequest(t *testing.T) {
	r := routeInfo{Method: "GET", Handler: "h"}
	ep := buildOneEndpoint("GET", "/users", r, "routes/u.ts", nil, []byte(""), map[string]*sitter.Node{})
	if ep.Request != nil {
		t.Fatalf("expected nil request, got %v", ep.Request)
	}
}
