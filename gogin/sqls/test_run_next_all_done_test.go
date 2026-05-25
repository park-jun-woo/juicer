//ff:func feature=ratchet type=session control=sequence
//ff:what TestRunNext_AllDone 테스트
package sqls

import (
	"testing"
)

func TestRunNext_AllDone(t *testing.T) {
	setupSessionDir(t)
	sess := &Session{
		RepoDir:    ".",
		QueriesDir: ".",
		Methods:    []MethodStatus{},
	}
	setupTestSession(t, sess)

	err := RunNext("", "")
	if err != nil {
		t.Fatalf("RunNext() error: %v", err)
	}
}
