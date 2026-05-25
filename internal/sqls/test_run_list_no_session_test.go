//ff:func feature=ratchet type=session control=sequence
//ff:what TestRunList_NoSession 테스트
package sqls

import (
	"testing"
)

func TestRunList_NoSession(t *testing.T) {
	setupSessionDir(t)
	err := RunList()
	if err != nil {
		t.Fatalf("RunList() error: %v", err)
	}
}
