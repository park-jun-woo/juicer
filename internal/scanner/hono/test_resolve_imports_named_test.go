//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestResolveImports_Named 테스트
package hono

import (
	"path/filepath"
	"testing"
)

func TestResolveImports_Named(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "schemas.ts", "export const s = 1\n")
	writeFile(t, dir, "app.ts", `import { s } from "./schemas"`+"\n")
	fi, err := parseFile(filepath.Join(dir, "app.ts"))
	if err != nil {
		t.Fatal(err)
	}
	imports := resolveImports(fi, dir)
	if imports["s"] != filepath.Join(dir, "schemas.ts") {
		t.Fatalf("got %v", imports)
	}
}
