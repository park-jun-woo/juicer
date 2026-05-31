//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what referentialAction onDelete/onUpdate 값 → SQL 절 매핑 테스트
package prisma

import "testing"

func TestReferentialAction(t *testing.T) {
	rel := "fields: [authorId], references: [id], onDelete: Cascade, onUpdate: SetNull"
	if got := referentialAction(rel, "onDelete"); got != "CASCADE" {
		t.Errorf("onDelete: got %q, want CASCADE", got)
	}
	if got := referentialAction(rel, "onUpdate"); got != "SET NULL" {
		t.Errorf("onUpdate: got %q, want SET NULL", got)
	}
	if got := referentialAction(rel, "onDelete2"); got != "" {
		t.Errorf("absent key: got %q", got)
	}
	if got := referentialAction("onDelete: Bogus", "onDelete"); got != "" {
		t.Errorf("unknown action: got %q", got)
	}
}
