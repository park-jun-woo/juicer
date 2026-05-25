//ff:func feature=ddl type=parse control=sequence
//ff:what TestStripLeadingComments_Basic 테스트
package ddl

import "testing"

func TestStripLeadingComments_Basic(t *testing.T) {
	got := stripLeadingComments("-- comment\nCREATE TABLE t (id INT)")
	if got != "CREATE TABLE t (id INT)" {
		t.Fatalf("got %q", got)
	}
}
