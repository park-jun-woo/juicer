//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestBuildOneEndpoint_WithRequest 테스트
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
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
