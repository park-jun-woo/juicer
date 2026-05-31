//ff:func feature=scan type=test topic=nestjs control=sequence
//ff:what namedDTOType DTO 베이스명/배열/유효성 판정 테스트
package nestjs

import "testing"

func TestNamedDTOType(t *testing.T) {
	base, arr, ok := namedDTOType("AlbumResponseDto[]")
	if !ok || base != "AlbumResponseDto" || !arr {
		t.Errorf("array dto: (%q,%v,%v)", base, arr, ok)
	}
	base, arr, ok = namedDTOType("UserDto")
	if !ok || base != "UserDto" || arr {
		t.Errorf("scalar dto: (%q,%v,%v)", base, arr, ok)
	}
	if _, _, ok := namedDTOType("string"); ok {
		t.Error("primitive should be not ok")
	}
	if _, _, ok := namedDTOType("A|B"); ok {
		t.Error("union should be not ok")
	}
}
