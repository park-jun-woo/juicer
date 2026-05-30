//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestBuildResponses_WithFields 테스트
package fastify

import "testing"

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
	fieldsByStatus := map[string]int{}
	for _, r := range resps {
		fieldsByStatus[r.Status] = len(r.Fields)
	}
	if n, ok := fieldsByStatus["200"]; !ok || n != 1 {
		t.Errorf("200 fields = %d (present=%v)", n, ok)
	}
	if n, ok := fieldsByStatus["404"]; !ok || n != 0 {
		t.Errorf("404 should have no fields, got %d (present=%v)", n, ok)
	}
}
