//ff:func feature=ratchet type=session control=sequence
//ff:what TestRunReset_NoSession 테스트
package sqls

import (
	"testing"
)

func TestRunReset_NoSession(t *testing.T) {
	setupSessionDir(t)
	err := RunReset()
	if err != nil {
		t.Fatalf("RunReset() error: %v", err)
	}
}
