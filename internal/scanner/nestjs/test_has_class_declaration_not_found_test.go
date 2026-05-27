//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestHasClassDeclaration_NotFound 존재하지 않는 클래스 선언 확인 테스트
package nestjs

import "testing"

func TestHasClassDeclaration_NotFound(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "dto.ts", `
export class OtherDto { name: string; }
`)
	if hasClassDeclaration(dir+"/dto.ts", "CreateUserDto") {
		t.Fatal("expected false for missing class")
	}
}
