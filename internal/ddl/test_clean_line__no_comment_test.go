//ff:func feature=ddl type=parse control=sequence
//ff:what TestCleanLine_NoComment 테스트
package ddl

import "testing"

func TestCleanLine_NoComment(t *testing.T) {
	got := cleanLine("  id INT  ")
	if got != "id INT" {
		t.Fatalf("got %q", got)
	}
}
