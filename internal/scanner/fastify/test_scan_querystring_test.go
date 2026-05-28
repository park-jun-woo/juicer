//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what querystring 스키마 추출 테스트
package fastify

import "testing"

func TestScan_Querystring(t *testing.T) {
	dir := t.TempDir()
	src := `
import Fastify from "fastify";
const app = Fastify();
app.get("/search", {
  schema: {
    querystring: {
      type: "object",
      properties: {
        q: { type: "string" },
        page: { type: "integer" },
        limit: { type: "integer" }
      }
    }
  }
}, searchHandler);
`
	writeFile(t, dir, "app.ts", src)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(result.Endpoints))
	}
	ep := result.Endpoints[0]
	if ep.Request == nil {
		t.Fatal("expected request to be non-nil")
	}
	if len(ep.Request.Query) != 3 {
		t.Fatalf("expected 3 query params, got %d", len(ep.Request.Query))
	}
	qMap := make(map[string]string)
	for _, q := range ep.Request.Query {
		qMap[q.Name] = q.Type
	}
	if qMap["q"] != "string" {
		t.Errorf("q.Type: want string, got %s", qMap["q"])
	}
	if qMap["page"] != "integer" {
		t.Errorf("page.Type: want integer, got %s", qMap["page"])
	}
}
