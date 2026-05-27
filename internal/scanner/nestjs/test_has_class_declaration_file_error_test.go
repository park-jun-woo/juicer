//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestHasClassDeclaration_FileError 존재하지 않는 파일 에러 확인 테스트
package nestjs

import "testing"

func TestHasClassDeclaration_FileError(t *testing.T) {
	if hasClassDeclaration("/nonexistent/dto.ts", "Dto") {
		t.Fatal("expected false for missing file")
	}
}
