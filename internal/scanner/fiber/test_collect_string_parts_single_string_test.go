//ff:func feature=scan type=test control=sequence
//ff:what TestCollectStringParts_SingleString 테스트
package fiber

import "testing"

func TestCollectStringParts_SingleString(t *testing.T) {
	if got := collectFor(t, `"hello"`); len(got) != 1 || got[0] != "hello" {
		t.Fatalf("got %v", got)
	}
}
