//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestBuildRequest_WithSchema 테스트
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestBuildRequest_WithSchema(t *testing.T) {
	src := `
import Fastify from "fastify";
const app = Fastify();
app.post("/u", {
  schema: { body: { type: "object", properties: { name: { type: "string" } } } }
}, h);
`
	fi := mustParse(t, []byte(src))
	instances := collectInstances(fi)
	routes := extractRoutes(fi, instances)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	req, has := buildRequest(routes[0], nil, fi.Src, map[string]*sitter.Node{})
	if !has || req.Body == nil || len(req.Body.Fields) != 1 {
		t.Fatalf("expected body from schema, got has=%v body=%v", has, req.Body)
	}
}
