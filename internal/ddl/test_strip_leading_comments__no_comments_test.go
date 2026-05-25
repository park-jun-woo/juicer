//ff:func feature=ddl type=parse control=sequence
//ff:what TestStripLeadingComments_NoComments 테스트
package ddl

import "testing"

func TestStripLeadingComments_NoComments(t *testing.T) {
	got := stripLeadingComments("CREATE TABLE t (id INT)")
	if got != "CREATE TABLE t (id INT)" {
		t.Fatalf("got %q", got)
	}
}
