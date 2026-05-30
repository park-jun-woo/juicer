//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestApplySchemaToRequest_All 테스트
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

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
