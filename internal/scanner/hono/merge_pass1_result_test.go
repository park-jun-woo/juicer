//ff:func feature=scan type=test control=sequence topic=hono
//ff:what mergePass1Result 테스트
package hono

import (
	"path/filepath"
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func newMergeMaps() (map[string]*fileInfo, map[string]map[string]bool, map[string]string, map[string]*sitter.Node, *[]routeGroup, map[string]map[string]string) {
	return map[string]*fileInfo{},
		map[string]map[string]bool{},
		map[string]string{},
		map[string]*sitter.Node{},
		&[]routeGroup{},
		map[string]map[string]string{}
}

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

func TestMergePass1Result_NilResult(t *testing.T) {
	// non-existent file -> scanOneFilePass1 returns nil -> early return, maps untouched
	parsed, vars, bp, schemas, groups, imports := newMergeMaps()
	mergePass1Result("/no/such/file.ts", "/no/such", parsed, vars, bp, schemas, groups, imports)
	if len(parsed) != 0 || len(vars) != 0 || len(bp) != 0 || len(*groups) != 0 {
		t.Fatal("expected no mutation for nil result")
	}
}

func TestMergePass1Result_NoVarsNoImports(t *testing.T) {
	// a file with no Hono vars / imports -> those maps stay empty (false branches)
	dir := t.TempDir()
	writeFile(t, dir, "plain.ts", "const x = 1\n")
	path := filepath.Join(dir, "plain.ts")
	parsed, vars, bp, schemas, groups, imports := newMergeMaps()
	mergePass1Result(path, dir, parsed, vars, bp, schemas, groups, imports)
	if parsed[path] == nil {
		t.Fatal("parsed should still be set")
	}
	if len(vars) != 0 {
		t.Fatal("expected no hono vars")
	}
	if len(imports) != 0 {
		t.Fatal("expected no imports")
	}
}
