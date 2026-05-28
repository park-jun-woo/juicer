//ff:func feature=scan type=test control=sequence topic=hono
//ff:what E2E 스캔 테스트: Zod validator → request body 필드 추출
package hono

import "testing"

func TestScan_ZodBody(t *testing.T) {
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
	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(result.Endpoints))
	}
	ep := result.Endpoints[0]
	if ep.Request == nil {
		t.Fatal("expected request")
	}
	if ep.Request.Body == nil {
		t.Fatal("expected body")
	}
	if len(ep.Request.Body.Fields) != 3 {
		t.Errorf("expected 3 body fields, got %d", len(ep.Request.Body.Fields))
	}
}
