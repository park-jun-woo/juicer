//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestResolveImports_UnresolvableRelative 테스트
package hono

import (
	"path/filepath"
	"testing"
)

func TestResolveImports_UnresolvableRelative(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "app.ts", `import { x } from "./missing"`+"\n")
	fi, _ := parseFile(filepath.Join(dir, "app.ts"))
	if imports := resolveImports(fi, dir); len(imports) != 0 {
		t.Fatalf("expected no imports, got %v", imports)
	}
}
