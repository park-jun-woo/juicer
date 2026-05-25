//ff:func feature=ratchet type=session control=sequence
//ff:what TestRunStatus_NoSession 테스트
package sqls

import (
	"testing"
)

func TestRunStatus_NoSession(t *testing.T) {
	setupSessionDir(t)
	err := RunStatus()
	if err != nil {
		t.Fatalf("RunStatus() error: %v", err)
	}
}
