//ff:func feature=ratchet type=session control=sequence
//ff:what TestCountStatus 테스트
package sqls

import (
	"testing"
)

func TestCountStatus(t *testing.T) {
	sess := &Session{
		Methods: []MethodStatus{
			{Status: "TODO"},
			{Status: "TODO"},
			{Status: "DONE"},
			{Status: "SKIP"},
		},
	}

	if countStatus(sess, "TODO") != 2 {
		t.Error("expected 2 TODO")
	}
	if countStatus(sess, "DONE") != 1 {
		t.Error("expected 1 DONE")
	}
	if countStatus(sess, "SKIP") != 1 {
		t.Error("expected 1 SKIP")
	}
}
