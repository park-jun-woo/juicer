//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestResolveImports_DefaultImport 테스트
package hono

import (
	"path/filepath"
	"testing"
)

func TestResolveImports_DefaultImport(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "mod.ts", "export default 1\n")
	writeFile(t, dir, "app.ts", `import mod from "./mod"`+"\n")
	fi, _ := parseFile(filepath.Join(dir, "app.ts"))
	imports := resolveImports(fi, dir)
	if imports["mod"] != filepath.Join(dir, "mod.ts") {
		t.Fatalf("got %v", imports)
	}
}
