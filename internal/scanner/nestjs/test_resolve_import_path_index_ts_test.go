//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveImportPath_IndexTS 테스트
package nestjs

import "testing"

func TestResolveImportPath_IndexTS(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "dto/index.ts", "export class Dto {}")
	result := resolveImportPath(dir, "./dto")
	if result == "" {
		t.Fatal("expected resolved path via index.ts")
	}
}
