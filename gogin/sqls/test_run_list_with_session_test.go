//ff:func feature=ratchet type=session control=sequence
//ff:what TestRunList_WithSession 테스트
package sqls

import (
	"testing"
)

func TestRunList_WithSession(t *testing.T) {
	setupSessionDir(t)
	sess := &Session{
		RepoDir:    ".",
		QueriesDir: ".",
		Methods: []MethodStatus{
			{ID: "A.B", Status: "TODO"},
			{ID: "A.C", Status: "DONE", QueryName: "GetAll"},
		},
	}
	setupTestSession(t, sess)

	err := RunList()
	if err != nil {
		t.Fatalf("RunList() error: %v", err)
	}
}
