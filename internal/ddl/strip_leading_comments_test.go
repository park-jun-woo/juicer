package ddl

import "testing"

func TestStripLeadingComments_Basic(t *testing.T) {
	got := stripLeadingComments("-- comment\nCREATE TABLE t (id INT)")
	if got != "CREATE TABLE t (id INT)" {
		t.Fatalf("got %q", got)
	}
}

func TestStripLeadingComments_NoComments(t *testing.T) {
	got := stripLeadingComments("CREATE TABLE t (id INT)")
	if got != "CREATE TABLE t (id INT)" {
		t.Fatalf("got %q", got)
	}
}

func TestStripLeadingComments_MultipleComments(t *testing.T) {
	got := stripLeadingComments("-- a\n-- b\n\nCREATE TABLE t (id INT)")
	if got != "CREATE TABLE t (id INT)" {
		t.Fatalf("got %q", got)
	}
}
