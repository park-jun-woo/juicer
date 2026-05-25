//ff:func feature=ratchet type=session control=sequence
//ff:what TestRunReset_WithSession 테스트
package sqls

import (
	"testing"
)

func TestRunReset_WithSession(t *testing.T) {
	setupSessionDir(t)
	sess := &Session{
		RepoDir:    ".",
		QueriesDir: ".",
		Methods:    []MethodStatus{},
	}
	setupTestSession(t, sess)

	err := RunReset()
	if err != nil {
		t.Fatalf("RunReset() error: %v", err)
	}

	if SessionExists() {
		t.Error("expected session to be deleted")
	}
}
