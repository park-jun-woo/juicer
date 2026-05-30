//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what applySchemaToRequest 테스트
package fastify

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

func schemaFromRoute(t *testing.T, routeSrc string) (*schemaInfo, []byte) {
	t.Helper()
	fi := mustParse(t, []byte(routeSrc))
	instances := collectInstances(fi)
	routes := extractRoutes(fi, instances)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	si := extractJSONSchema(routes[0].Schema, fi.Src)
	if si == nil {
		t.Fatal("nil schemaInfo")
	}
	return si, fi.Src
}

func TestApplySchemaToRequest_All(t *testing.T) {
	src := `
import Fastify from "fastify";
const app = Fastify();
app.post("/u/:id", {
  schema: {
    params: { type: "object", properties: { id: { type: "integer" } } },
    querystring: { type: "object", properties: { q: { type: "string" } } },
    body: { type: "object", properties: { name: { type: "string" } } }
  }
}, h);
`
	si, fiSrc := schemaFromRoute(t, src)
	req := &scanner.Request{}
	hasReq := false
	applySchemaToRequest(si, fiSrc, req, &hasReq, map[string]*sitter.Node{})

	if !hasReq {
		t.Fatal("hasReq should be true")
	}
	if len(req.PathParams) != 1 || req.PathParams[0].Name != "id" {
		t.Errorf("path params: %v", req.PathParams)
	}
	if len(req.Query) != 1 || req.Query[0].Name != "q" {
		t.Errorf("query: %v", req.Query)
	}
	if req.Body == nil || len(req.Body.Fields) != 1 || req.Body.Fields[0].Name != "name" {
		t.Errorf("body: %v", req.Body)
	}
}

func TestApplySchemaToRequest_Empty(t *testing.T) {
	// schemaInfo with all nil sections -> hasReq stays false
	si := &schemaInfo{}
	req := &scanner.Request{}
	hasReq := false
	applySchemaToRequest(si, []byte(""), req, &hasReq, map[string]*sitter.Node{})
	if hasReq || req.Body != nil || len(req.PathParams) != 0 || len(req.Query) != 0 {
		t.Fatalf("expected no changes, got %v hasReq=%v", req, hasReq)
	}
}
