//ff:func feature=sql type=parse control=sequence
//ff:what TestSessionPath_ReturnsPath 테스트
package sqls

import "testing"

func TestSessionPath_ReturnsPath(t *testing.T) {
	got := sessionPath()
	if got == "" {
		t.Fatal("expected non-empty")
	}
}
