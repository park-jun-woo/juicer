//ff:func feature=scan type=test control=sequence topic=hono
//ff:what zValidator() 미들웨어에서 target + schema 추출 테스트
package hono

import "testing"

func TestExtractZodValidators(t *testing.T) {
	src := []byte(`
import { Hono } from "hono"
import { zValidator } from "@hono/zod-validator"
const app = new Hono()
app.post("/users", zValidator("json", createUserSchema), (c) => {
  return c.json({ ok: true })
})
`)
	fi := mustParse(t, src)
	vars := collectHonoVars(fi)
	routes := collectRoutes(fi, vars)
	if len(routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(routes))
	}
	if len(routes[0].ZodValidators) != 1 {
		t.Fatalf("expected 1 zod validator, got %d", len(routes[0].ZodValidators))
	}
	v := routes[0].ZodValidators[0]
	if v.Target != "json" {
		t.Errorf("expected target json, got %s", v.Target)
	}
	if v.SchemaName != "createUserSchema" {
		t.Errorf("expected schemaName createUserSchema, got %s", v.SchemaName)
	}
}
