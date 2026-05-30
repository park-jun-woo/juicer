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

func TestExtractZodValidators_MixedArgs(t *testing.T) {
	// some args are zValidator calls, some are not (handler / non-validator)
	fi := mustParse(t, []byte(`app.post("/x", zValidator("json", s), other(), handler);`+"\n"))
	args := findAllByType(fi.Root, "arguments")[0]
	nodes := collectArgNodes(args)
	vs := extractZodValidators(nodes, fi.Src)
	if len(vs) != 1 || vs[0].Target != "json" || vs[0].SchemaName != "s" {
		t.Fatalf("got %+v", vs)
	}
}

func TestExtractZodValidators_None(t *testing.T) {
	fi := mustParse(t, []byte(`app.get("/x", handler);`+"\n"))
	args := findAllByType(fi.Root, "arguments")[0]
	nodes := collectArgNodes(args)
	if vs := extractZodValidators(nodes, fi.Src); len(vs) != 0 {
		t.Fatalf("expected none, got %+v", vs)
	}
}
