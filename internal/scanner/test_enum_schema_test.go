//ff:func feature=scan type=test control=sequence
//ff:what EnumSchema 문자열 enum OpenAPI 스키마 객체 생성 테스트
package scanner

import (
	"reflect"
	"testing"
)

func TestEnumSchema(t *testing.T) {
	got := EnumSchema([]string{"A", "B"})
	if got["type"] != "string" {
		t.Errorf("type: %v", got["type"])
	}
	if !reflect.DeepEqual(got["enum"], []string{"A", "B"}) {
		t.Errorf("enum: %v", got["enum"])
	}
}
