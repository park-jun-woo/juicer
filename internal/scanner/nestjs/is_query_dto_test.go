//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestIsQueryDTO 테스트
package nestjs

import "testing"

func TestIsQueryDTO(t *testing.T) {
	if !isQueryDTO("", "ListUserReqDto") {
		t.Fatal("expected true for empty arg + DTO type")
	}
	if isQueryDTO("page", "string") {
		t.Fatal("expected false for named arg")
	}
	if isQueryDTO("", "string") {
		t.Fatal("expected false for primitive type")
	}
	if isQueryDTO("", "number") {
		t.Fatal("expected false for number type")
	}
}
