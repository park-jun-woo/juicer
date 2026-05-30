//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestBuildOneEndpoint_NoRequest 테스트
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestBuildOneEndpoint_NoRequest(t *testing.T) {
	r := routeInfo{Method: "GET", Handler: "h"}
	ep := buildOneEndpoint("GET", "/users", r, "routes/u.ts", nil, []byte(""), map[string]*sitter.Node{})
	if ep.Request != nil {
		t.Fatalf("expected nil request, got %v", ep.Request)
	}
}
