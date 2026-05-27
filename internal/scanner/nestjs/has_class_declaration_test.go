//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestHasClassDeclaration_Found 클래스 선언 존재 확인 테스트
package nestjs

import "testing"

func TestHasClassDeclaration_Found(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "dto.ts", `
export class CreateUserDto {
  name: string;
}
`)
	if !hasClassDeclaration(dir+"/dto.ts", "CreateUserDto") {
		t.Fatal("expected true for existing class")
	}
}
