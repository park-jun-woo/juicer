//ff:func feature=ratchet type=session control=sequence
//ff:what TestCreateSession_NoQueries 테스트
package sqls

import (
	"testing"
)

func TestCreateSession_NoQueries(t *testing.T) {
	setupSessionDir(t)
	err := createSession("somedir", "")
	if err == nil {
		t.Error("expected error for missing --queries")
	}
}
