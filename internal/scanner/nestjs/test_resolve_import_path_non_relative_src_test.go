//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveImportPath_NonRelativeSrc 테스트
package nestjs

import "testing"

func TestResolveImportPath_NonRelativeSrc(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/users/dto/create-user.dto.ts", `
export class CreateUserDto {
  name: string;
}
`)
	result := resolveImportPath("/some/other/dir", "src/users/dto/create-user.dto", dir)
	if result == "" {
		t.Fatal("expected resolved path for non-relative src/ import")
	}
}
