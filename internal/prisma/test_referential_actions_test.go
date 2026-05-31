//ff:func feature=prisma type=test control=sequence topic=prisma
//ff:what referentialActions ON DELETE/ON UPDATE 절 렌더링 테스트
package prisma

import "testing"

func TestReferentialActions(t *testing.T) {
	rel := "onDelete: Cascade, onUpdate: Restrict"
	if got := referentialActions(rel); got != " ON DELETE CASCADE ON UPDATE RESTRICT" {
		t.Errorf("both: got %q", got)
	}
	if got := referentialActions("onDelete: SetNull"); got != " ON DELETE SET NULL" {
		t.Errorf("delete only: got %q", got)
	}
	if got := referentialActions(""); got != "" {
		t.Errorf("none: got %q", got)
	}
}
