//ff:func feature=ratchet type=session control=sequence
//ff:what TestRunNext_NoSession_NoFlags 테스트
package sqls

import (
	"testing"
)

func TestRunNext_NoSession_NoFlags(t *testing.T) {
	setupSessionDir(t)
	err := RunNext("", "")
	if err == nil {
		t.Error("expected error for missing --repo")
	}
}
