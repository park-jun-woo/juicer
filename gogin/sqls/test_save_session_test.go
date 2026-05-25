//ff:func feature=sql type=session control=sequence
//ff:what TestSaveSession 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSaveSession(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	sess := &Session{
		RepoDir:    "repo",
		QueriesDir: "queries",
		Methods: []MethodStatus{
			{ID: "A.B", Status: "TODO"},
		},
	}

	err := SaveSession(sess)
	if err != nil {
		t.Fatalf("SaveSession() error: %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(filepath.Join(".huma", "sql-session.json")); err != nil {
		t.Errorf("session file not found: %v", err)
	}

	// Load and verify
	loaded, err := LoadSession()
	if err != nil {
		t.Fatalf("LoadSession() error: %v", err)
	}
	if loaded.RepoDir != "repo" {
		t.Errorf("expected RepoDir 'repo', got %q", loaded.RepoDir)
	}
	if len(loaded.Methods) != 1 {
		t.Errorf("expected 1 method, got %d", len(loaded.Methods))
	}
}
