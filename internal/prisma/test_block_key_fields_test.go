//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what blockKeyFields 접두 블록 속성의 필드 목록 추출 테스트
package prisma

import (
	"reflect"
	"testing"
)

func TestBlockKeyFields(t *testing.T) {
	attrs := []string{"@@unique([orgId, email])", "@@id([a, b])"}
	got, ok := blockKeyFields(attrs, "@@id")
	if !ok || !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Errorf("got (%v,%v)", got, ok)
	}
	if _, ok := blockKeyFields(attrs, "@@index"); ok {
		t.Error("absent prefix must be false")
	}
}
