//ff:func feature=ratchet type=session control=sequence
//ff:what TestRunStatus_WithSession 테스트
package sqls

import (
	"testing"
)

func TestRunStatus_WithSession(t *testing.T) {
	setupSessionDir(t)
	sess := &Session{
		RepoDir:    ".",
		QueriesDir: ".",
		Methods: []MethodStatus{
			{ID: "A.B", Status: "TODO"},
			{ID: "A.C", Status: "DONE"},
			{ID: "A.D", Status: "SKIP"},
		},
	}
	setupTestSession(t, sess)

	err := RunStatus()
	if err != nil {
		t.Fatalf("RunStatus() error: %v", err)
	}
}
