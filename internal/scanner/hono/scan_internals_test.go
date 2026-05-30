//ff:func feature=scan type=test control=sequence topic=hono
//ff:what scanOneFilePass1 / scanPass1 / scanPass2 직접 테스트
package hono

import (
	"path/filepath"
	"testing"
)

const scanInternalsSrc = `
import { Hono } from "hono"
const app = new Hono()
app.get("/users/:id", (c) => c.json({ id: 1 }))
app.post("/users", (c) => c.json({ ok: true }, 201))
`

func TestScanOneFilePass1_OK(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app.ts", scanInternalsSrc)
	r := scanOneFilePass1(filepath.Join(dir, "app.ts"), dir)
	if r == nil {
		t.Fatal("nil result")
	}
	if r.fi == nil || len(r.vars) == 0 {
		t.Fatalf("missing fi/vars: %+v", r)
	}
}

func TestScanOneFilePass1_ParseError(t *testing.T) {
	// non-existent file -> parseFile errors -> nil
	if r := scanOneFilePass1("/no/such.ts", "/no"); r != nil {
		t.Fatal("expected nil for parse error")
	}
}

func TestScanPass1(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app.ts", scanInternalsSrc)
	ctx := scanPass1([]string{filepath.Join(dir, "app.ts")}, dir)
	if ctx == nil || len(ctx.parsed) != 1 {
		t.Fatalf("unexpected ctx: %+v", ctx)
	}
	if ctx.absRoot != dir {
		t.Fatalf("absRoot: %s", ctx.absRoot)
	}
}

func TestScanPass2(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app.ts", scanInternalsSrc)
	ctx := scanPass1([]string{filepath.Join(dir, "app.ts")}, dir)
	eps := scanPass2(ctx)
	if len(eps) != 2 {
		t.Fatalf("expected 2 endpoints, got %d: %+v", len(eps), eps)
	}
}

func TestScanPass2_SkipFilesWithoutVars(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app.ts", scanInternalsSrc)
	writeFile(t, dir, "plain.ts", "const x = 1\n")
	ctx := scanPass1([]string{
		filepath.Join(dir, "app.ts"),
		filepath.Join(dir, "plain.ts"),
	}, dir)
	eps := scanPass2(ctx)
	// plain.ts has no hono vars -> skipped, still 2 endpoints
	if len(eps) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(eps))
	}
}
