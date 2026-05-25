//ff:func feature=ddl type=parse control=sequence
//ff:what TestStripInlineComments_Basic 테스트
package ddl

import "testing"

func TestStripInlineComments_Basic(t *testing.T) {
	got := stripInlineComments("id INT -- primary key\nname TEXT -- user name")
	if got != "id INT \nname TEXT " {
		t.Fatalf("got %q", got)
	}
}
