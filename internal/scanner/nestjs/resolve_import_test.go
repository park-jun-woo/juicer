//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveImportPath_TSExtension 테스트
package nestjs

import "testing"

func TestResolveImportPath_TSExtension(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "dto/create-user.dto.ts", "export class Dto {}")
	result := resolveImportPath(dir, "./dto/create-user.dto")
	if result == "" {
		t.Fatal("expected resolved path")
	}
}
