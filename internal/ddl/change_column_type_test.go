//ff:func feature=ddl type=test control=sequence
//ff:what changeColumnType 단일 토큰 입력 시 원본 반환 테스트
package ddl

import "testing"

func TestChangeColumnType_ShortRaw(t *testing.T) {
	got := changeColumnType("id", "BIGINT")
	if got != "id" {
		t.Errorf("changeColumnType with single token should return raw unchanged, got %q", got)
	}
}
