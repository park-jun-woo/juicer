//ff:func feature=ddl type=parse control=sequence
//ff:what TestApplyStatement_LeadingComment 테스트
package ddl

import "testing"

func TestApplyStatement_LeadingComment(t *testing.T) {
	tables := make(map[string]*Table)
	applyStatement(tables, "-- comment\nCREATE TABLE items (id INT)")
	if tables["items"] == nil {
		t.Fatal("expected items table")
	}
}
