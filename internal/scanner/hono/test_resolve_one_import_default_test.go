//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestResolveOneImport_Default 테스트
package hono

import (
	"path/filepath"
	"testing"
)

func TestResolveOneImport_Default(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "mod.ts", "export default 1\n")
	imp := resolveOne(t, dir, `import mod from "./mod"`)
	if imp["mod"] != filepath.Join(dir, "mod.ts") {
		t.Fatalf("got %v", imp)
	}
}
