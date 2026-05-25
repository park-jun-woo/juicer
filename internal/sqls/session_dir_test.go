//ff:func feature=sql type=parse control=sequence
//ff:what TestSessionDir_ReturnsDir 테스트
package sqls

import "testing"

func TestSessionDir_ReturnsDir(t *testing.T) {
	got := SessionDir()
	if got != ".huma" {
		t.Fatalf("expected .huma, got %s", got)
	}
}
