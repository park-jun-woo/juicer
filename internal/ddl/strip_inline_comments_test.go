package ddl

import "testing"

func TestStripInlineComments_Basic(t *testing.T) {
	got := stripInlineComments("id INT -- primary key\nname TEXT -- user name")
	if got != "id INT \nname TEXT " {
		t.Fatalf("got %q", got)
	}
}

func TestStripInlineComments_NoComments(t *testing.T) {
	got := stripInlineComments("id INT\nname TEXT")
	if got != "id INT\nname TEXT" {
		t.Fatalf("got %q", got)
	}
}
