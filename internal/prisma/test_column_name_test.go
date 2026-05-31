//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what columnName @map 오버라이드 및 기본 필드명 테스트
package prisma

import "testing"

func TestColumnName(t *testing.T) {
	if got := columnName(field{name: "userId"}); got != "userId" {
		t.Errorf("no @map: got %q, want userId", got)
	}
	mapped := field{name: "userId", attrs: []string{`@map("user_id")`}}
	if got := columnName(mapped); got != "user_id" {
		t.Errorf("@map: got %q, want user_id", got)
	}
	other := field{name: "x", attrs: []string{"@id", `@map("col")`}}
	if got := columnName(other); got != "col" {
		t.Errorf("@map among others: got %q, want col", got)
	}
}
