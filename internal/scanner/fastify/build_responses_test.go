//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what buildResponses 테스트
package fastify

import "testing"

func TestBuildResponses_NilSchema(t *testing.T) {
	if r := buildResponses(routeInfo{}, []byte("")); r != nil {
		t.Fatalf("expected nil for nil schema, got %v", r)
	}
}

func TestBuildResponses_NoSchemaKey(t *testing.T) {
	obj, src := firstObject(t, `{ config: {} }`)
	if r := buildResponses(routeInfo{Schema: obj}, src); r != nil {
		t.Fatalf("expected nil when schema info nil, got %v", r)
	}
}

func TestBuildResponses_WithFields(t *testing.T) {
	src := `
import Fastify from "fastify";
const app = Fastify();
app.get("/u", {
  schema: {
    response: {
      "200": { type: "object", properties: { id: { type: "integer" } } },
      "404": { type: "object" }
    }
  }
}, h);
`
	fi := mustParse(t, []byte(src))
	routes := extractRoutes(fi, collectInstances(fi))
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	resps := buildResponses(routes[0], fi.Src)
	if len(resps) != 2 {
		t.Fatalf("expected 2 responses, got %d", len(resps))
	}
	// 200 has fields, 404 has none
	var got200, got404 bool
	for _, r := range resps {
		if r.Status == "200" {
			got200 = true
			if len(r.Fields) != 1 {
				t.Errorf("200 fields = %d", len(r.Fields))
			}
		}
		if r.Status == "404" {
			got404 = true
			if len(r.Fields) != 0 {
				t.Errorf("404 should have no fields, got %d", len(r.Fields))
			}
		}
	}
	if !got200 || !got404 {
		t.Errorf("missing status: 200=%v 404=%v", got200, got404)
	}
}
