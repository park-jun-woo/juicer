//ff:func feature=sql type=session control=sequence
//ff:what TestSessionDir 테스트
package sqls

import (
	"testing"
)

func TestSessionDir(t *testing.T) {
	got := SessionDir()
	if got != ".codist" {
		t.Errorf("expected '.codist', got %q", got)
	}
}
