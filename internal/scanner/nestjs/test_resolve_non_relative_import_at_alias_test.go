//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveNonRelativeImport_AtAlias 테스트
package nestjs

import (
	"path/filepath"
	"testing"
)

func TestResolveNonRelativeImport_AtAlias(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/decorators/field.ts", "x")
	got := resolveNonRelativeImport(dir, "@/decorators/field")
	if got != filepath.Join(dir, "src/decorators/field.ts") {
		t.Fatalf("got %q", got)
	}
}
