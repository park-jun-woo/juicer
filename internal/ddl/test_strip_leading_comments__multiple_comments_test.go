//ff:func feature=ddl type=parse control=sequence
//ff:what TestStripLeadingComments_MultipleComments 테스트
package ddl

import "testing"

func TestStripLeadingComments_MultipleComments(t *testing.T) {
	got := stripLeadingComments("-- a\n-- b\n\nCREATE TABLE t (id INT)")
	if got != "CREATE TABLE t (id INT)" {
		t.Fatalf("got %q", got)
	}
}
