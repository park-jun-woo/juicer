//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestMergePass1Result_NoVarsNoImports 테스트
package hono

import (
	"path/filepath"
	"testing"
)

func TestMergePass1Result_NoVarsNoImports(t *testing.T) {

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
