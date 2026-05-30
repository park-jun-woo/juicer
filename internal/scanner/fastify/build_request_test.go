//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what buildRequest 테스트
package fastify

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestBuildRequest_PathParamsOnly(t *testing.T) {
	r := routeInfo{Method: "GET"}
	req, has := buildRequest(r, []string{"id"}, []byte(""), map[string]*sitter.Node{})
	if !has || len(req.PathParams) != 1 {
		t.Fatalf("expected path params, got has=%v %v", has, req.PathParams)
	}
}

func TestBuildRequest_NoSchemaNoParams(t *testing.T) {
	r := routeInfo{Method: "GET", Schema: nil}
	req, has := buildRequest(r, nil, []byte(""), map[string]*sitter.Node{})
	if has || req.Body != nil {
		t.Fatalf("expected empty request, got has=%v", has)
	}
}

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

func TestBuildRequest_SchemaPresentButNoSchemaKey(t *testing.T) {
	// Schema is an object node but has no "schema" key -> extractJSONSchema nil.
	obj, src := firstObject(t, `{ config: { rateLimit: true } }`)
	r := routeInfo{Method: "GET", Schema: obj}
	req, has := buildRequest(r, nil, src, map[string]*sitter.Node{})
	if has || req.Body != nil {
		t.Fatalf("expected empty request when schema info nil, got has=%v", has)
	}
}
