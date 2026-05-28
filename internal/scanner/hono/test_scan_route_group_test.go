//ff:func feature=scan type=test control=iteration dimension=1 topic=hono
//ff:what app.route() 그룹 prefix 전파 E2E 테스트
package hono

import "testing"

func TestScan_RouteGroup(t *testing.T) {
	dir := t.TempDir()

	appSrc := `
import { Hono } from "hono"

const app = new Hono()
const users = new Hono()

users.get("/", (c) => c.json([]))
users.get("/:id", (c) => c.json({}))
users.post("/", (c) => c.json({}, 201))

app.route("/api/users", users)
app.get("/health", (c) => c.json({ ok: true }))
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
		"POST /api/users",
		"GET /health",
	}
	for _, e := range expected {
		if !found[e] {
			t.Errorf("missing endpoint %s, got %v", e, found)
		}
	}
}
