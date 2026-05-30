//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestResolveOneImport_Named 테스트
package hono

import (
	"path/filepath"
	"testing"
)

func TestResolveOneImport_Named(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "schemas.ts", "export const a = 1\n")
	imp := resolveOne(t, dir, `import { a } from "./schemas"`)
	if imp["a"] != filepath.Join(dir, "schemas.ts") {
		t.Fatalf("got %v", imp)
	}
}
