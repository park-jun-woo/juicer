//ff:func feature=ratchet type=session control=sequence
//ff:what TestRunSkip_WithSession 테스트
package sqls

import (
	"testing"
)

func TestRunSkip_WithSession(t *testing.T) {
	setupSessionDir(t)
	sess := &Session{
		RepoDir:    ".",
		QueriesDir: ".",
		Methods: []MethodStatus{
			{ID: "A.B", Status: "TODO"},
		},
	}
	setupTestSession(t, sess)

	err := RunSkip()
	if err != nil {
		t.Fatalf("RunSkip() error: %v", err)
	}

	// Verify it's now SKIP
	updated, _ := LoadSession()
	if updated.Methods[0].Status != "SKIP" {
		t.Errorf("expected SKIP, got %q", updated.Methods[0].Status)
	}
}
