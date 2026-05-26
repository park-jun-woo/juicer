//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveImportPath_AtAlias 테스트
package nestjs

import "testing"

func TestResolveImportPath_AtAlias(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/decorators/field.ts", `export function Field() {}`)
	result := resolveImportPath("/some/other/dir", "@/decorators/field", dir)
	if result == "" {
		t.Fatal("expected resolved path for @/ alias import")
	}
}
