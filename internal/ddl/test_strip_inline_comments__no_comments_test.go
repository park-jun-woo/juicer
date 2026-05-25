//ff:func feature=ddl type=parse control=sequence
//ff:what TestStripInlineComments_NoComments 테스트
package ddl

import "testing"

func TestStripInlineComments_NoComments(t *testing.T) {
	got := stripInlineComments("id INT\nname TEXT")
	if got != "id INT\nname TEXT" {
		t.Fatalf("got %q", got)
	}
}
