//ff:func feature=scan type=test control=sequence topic=hono
//ff:what buildEndpointsFromRoute 테스트
package hono

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func minimalCtx() *scanContext {
	return &scanContext{
		parsed:    map[string]*fileInfo{},
		honoVars:  map[string]map[string]bool{},
		basePaths: map[string]string{},
		schemas:   map[string]*sitter.Node{},
		prefixMap: map[string]string{},
		imports:   map[string]map[string]string{},
	}
}

func TestBuildEndpointsFromRoute_Single(t *testing.T) {
	r := routeInfo{Method: "GET", Path: "/users/:id", Handler: "h", OwnerVar: "app"}
	ctx := minimalCtx()
	ctx.prefixMap[prefixKey("/abs/app.ts", "app")] = "/api"
	fi := &fileInfo{Path: "/abs/app.ts"}

	eps := buildEndpointsFromRoute(r, map[string]bool{}, ctx, fi, "app.ts", "/abs/app.ts")
	if len(eps) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(eps))
	}
	if eps[0].Method != "GET" || eps[0].Path != "/api/users/{id}" {
		t.Fatalf("ep = %+v", eps[0])
	}
}

func TestBuildEndpointsFromRoute_All(t *testing.T) {
	r := routeInfo{Method: "all", Path: "/", Handler: "h"}
	ctx := minimalCtx()
	fi := &fileInfo{Path: "/abs/x.ts"}
	eps := buildEndpointsFromRoute(r, map[string]bool{}, ctx, fi, "x.ts", "/abs/x.ts")
	if len(eps) != 5 {
		t.Fatalf("expected 5 endpoints for 'all', got %d", len(eps))
	}
}
