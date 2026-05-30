//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestResolveImports_ExternalSkipped 테스트
package hono

import (
	"path/filepath"
	"testing"
)

func TestResolveImports_ExternalSkipped(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app.ts", `import { Hono } from "hono"`+"\n")
	fi, _ := parseFile(filepath.Join(dir, "app.ts"))
	if imports := resolveImports(fi, dir); len(imports) != 0 {
		t.Fatalf("expected no imports, got %v", imports)
	}
}
