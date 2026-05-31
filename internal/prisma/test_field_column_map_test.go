//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what fieldColumnMap 필드명 → 컬럼명(@map 반영) 매핑 테스트
package prisma

import "testing"

func TestFieldColumnMap(t *testing.T) {
	m := model{fields: []field{
		{name: "id"},
		{name: "userId", attrs: []string{`@map("user_id")`}},
	}}
	got := fieldColumnMap(m)
	if got["id"] != "id" || got["userId"] != "user_id" {
		t.Errorf("got %v", got)
	}
	if len(got) != 2 {
		t.Errorf("len = %d, want 2", len(got))
	}
}
