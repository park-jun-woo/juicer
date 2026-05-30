//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestMergePass1Result_Full 테스트
package hono

import (
	"path/filepath"
	"testing"
)

func TestMergePass1Result_Full(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "other.ts", "export const helper = 1\n")
	src := `
import { Hono } from "hono"
import { z } from "zod"
import { helper } from "./other"
const app = new Hono().basePath("/api")
const sub = new Hono()
const schema = z.object({ a: z.string() })
app.route("/sub", sub)
`
	writeFile(t, dir, "app.ts", src)
	path := filepath.Join(dir, "app.ts")

	parsed, vars, bp, schemas, groups, imports := newMergeMaps()
	mergePass1Result(path, dir, parsed, vars, bp, schemas, groups, imports)

	if parsed[path] == nil {
		t.Fatal("parsed not populated")
	}
	if len(vars[path]) == 0 {
		t.Fatal("hono vars not populated")
	}
	if len(bp) == 0 {
		t.Fatal("basePaths not populated")
	}
	if len(schemas) == 0 {
		t.Fatal("schemas not populated")
	}
	if len(*groups) == 0 {
		t.Fatal("groups not populated")
	}
	if len(imports[path]) == 0 {
		t.Fatal("imports not populated")
	}
}
