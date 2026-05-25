//ff:func feature=ddl type=parse control=sequence
//ff:what TestCleanLine_Empty 테스트
package ddl

import "testing"

func TestCleanLine_Empty(t *testing.T) {
	got := cleanLine("")
	if got != "" {
		t.Fatalf("got %q", got)
	}
}
