//ff:func feature=scan type=test control=iteration dimension=1 topic=hono
//ff:what basePath prefix 전파 E2E 테스트
package hono

import "testing"

func TestScan_BasePath(t *testing.T) {
	dir := t.TempDir()

	appSrc := `
import { Hono } from "hono"
const app = new Hono().basePath("/api")
app.get("/users", (c) => c.json([]))
app.get("/users/:id", (c) => c.json({}))
`
	writeFile(t, dir, "app.ts", appSrc)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}

	found := map[string]bool{}
	for _, ep := range result.Endpoints {
		found[ep.Method+" "+ep.Path] = true
	}
	expected := []string{
		"GET /api/users",
		"GET /api/users/{id}",
	}
	for _, e := range expected {
		if !found[e] {
			t.Errorf("missing endpoint %s, got %v", e, found)
		}
	}
}
