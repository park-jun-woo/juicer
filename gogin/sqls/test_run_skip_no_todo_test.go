//ff:func feature=ratchet type=session control=sequence
//ff:what TestRunSkip_NoTODO 테스트
package sqls

import (
	"testing"
)

func TestRunSkip_NoTODO(t *testing.T) {
	setupSessionDir(t)
	sess := &Session{
		RepoDir:    ".",
		QueriesDir: ".",
		Methods: []MethodStatus{
			{ID: "A.B", Status: "DONE"},
		},
	}
	setupTestSession(t, sess)

	err := RunSkip()
	if err != nil {
		t.Fatalf("RunSkip() error: %v", err)
	}
}
