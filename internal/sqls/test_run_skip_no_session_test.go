//ff:func feature=ratchet type=session control=sequence
//ff:what TestRunSkip_NoSession 테스트
package sqls

import (
	"testing"
)

func TestRunSkip_NoSession(t *testing.T) {
	setupSessionDir(t)
	err := RunSkip()
	if err != nil {
		t.Fatalf("RunSkip() error: %v", err)
	}
}
