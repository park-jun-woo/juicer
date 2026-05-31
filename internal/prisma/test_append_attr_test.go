//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what appendAttr 시작 인덱스 유효성 및 트림 테스트
package prisma

import (
	"reflect"
	"testing"
)

func TestAppendAttr(t *testing.T) {
	// start < 0 -> unchanged
	if got := appendAttr(nil, "x", -1, 1); got != nil {
		t.Errorf("start<0 must be unchanged: %v", got)
	}
	// trimmed non-empty -> appended
	got := appendAttr(nil, "  @id  ", 0, 7)
	if !reflect.DeepEqual(got, []string{"@id"}) {
		t.Errorf("got %v, want [@id]", got)
	}
	// whitespace-only -> not appended
	if got := appendAttr(nil, "    ", 0, 4); got != nil {
		t.Errorf("whitespace must not append: %v", got)
	}
}
