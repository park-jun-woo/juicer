//ff:func feature=scan type=test control=iteration dimension=1 topic=hono
//ff:what E2E 스캔 테스트: 기본 라우트 + 경로 파라미터
package hono

import "testing"

func TestScan_E2E(t *testing.T) {
	dir := t.TempDir()

	appSrc := `
import { Hono } from "hono"
import { zValidator } from "@hono/zod-validator"
import { z } from "zod"

const app = new Hono()

const createUserSchema = z.object({
  name: z.string().min(1),
  email: z.string().email(),
  age: z.number().int().optional()
})

app.get("/health", (c) => c.json({ ok: true }))
app.get("/users/:id", (c) => {
  const id = c.req.param("id")
  return c.json({ id })
})
app.post("/users", zValidator("json", createUserSchema), (c) => {
  const data = c.req.valid("json")
  return c.json(data, 201)
})
`
	writeFile(t, dir, "app.ts", appSrc)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) < 3 {
		t.Fatalf("expected at least 3 endpoints, got %d", len(result.Endpoints))
	}

	found := map[string]bool{}
	for _, ep := range result.Endpoints {
		found[ep.Method+" "+ep.Path] = true
	}
	expected := []string{
		"GET /health",
		"GET /users/{id}",
		"POST /users",
	}
	for _, e := range expected {
		if !found[e] {
			t.Errorf("missing endpoint %s, got %v", e, found)
		}
	}
}
