//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what primaryKeyConstraint 필드 @id / 블록 @@id / 없음 테스트
package prisma

import "testing"

func TestPrimaryKeyConstraint(t *testing.T) {
	s := schema{columns: map[string]map[string]string{"M": {"a": "a", "b": "b"}}}
	// field @id
	m := model{name: "M", fields: []field{{name: "id", attrs: []string{"@id"}}}}
	if got := primaryKeyConstraint(m, s); got != `PRIMARY KEY ("id")` {
		t.Errorf("field @id: got %q", got)
	}
	// block @@id
	m = model{name: "M", blockAttrs: []string{"@@id([a, b])"}}
	if got := primaryKeyConstraint(m, s); got != `PRIMARY KEY ("a", "b")` {
		t.Errorf("@@id: got %q", got)
	}
	// none
	if got := primaryKeyConstraint(model{name: "M"}, s); got != "" {
		t.Errorf("none: got %q", got)
	}
}
