//ff:func feature=ddl type=test control=sequence
//ff:what TestApplyStatement_LeadingCommentCov 테스트
package ddl

import "testing"

func TestApplyStatement_LeadingCommentCov(t *testing.T) {
	tables := make(map[string]*Table)
	applyStatement(tables, "-- comment\nCREATE TABLE foo (id INT)")
	if tables["foo"] == nil {
		t.Fatal("expected foo table")
	}
}
