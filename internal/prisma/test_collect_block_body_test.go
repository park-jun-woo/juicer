//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what collectBlockBody 닫는 중괄호까지 비빈 라인 수집 테스트
package prisma

import (
	"reflect"
	"testing"
)

func TestCollectBlockBody(t *testing.T) {
	lines := []string{"model User {", "id Int", "", "email String", "}", "next"}
	body, next := collectBlockBody(lines, 1)
	if !reflect.DeepEqual(body, []string{"id Int", "email String"}) {
		t.Errorf("body: %v", body)
	}
	if next != 5 {
		t.Errorf("next = %d, want 5", next)
	}
	// unterminated: returns all body and end index
	body, next = collectBlockBody([]string{"a", "b"}, 0)
	if len(body) != 2 || next != 2 {
		t.Errorf("unterminated: body=%v next=%d", body, next)
	}
}
