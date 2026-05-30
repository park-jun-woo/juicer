//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestBuildEndpointsFromRoute_All 테스트
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestBuildEndpointsFromRoute_All(t *testing.T) {
	r := routeInfo{Method: "all", Path: "/", Handler: "h"}
	eps := buildEndpointsFromRoute(r, "", "routes/x.ts", []byte(""), map[string]*sitter.Node{})
	if len(eps) != 5 {
		t.Fatalf("expected 5 endpoints for 'all', got %d", len(eps))
	}
}
