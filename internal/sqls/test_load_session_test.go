//ff:func feature=sql type=session control=sequence
//ff:what TestLoadSession 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadSession(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	t.Run("not found", func(t *testing.T) {
		_, err := LoadSession()
		if err == nil {
			t.Error("expected error for missing session")
		}
	})

	t.Run("invalid json", func(t *testing.T) {
		os.MkdirAll(".codist", 0o755)
		os.WriteFile(filepath.Join(".codist", "sql-session.json"), []byte("bad json"), 0o644)
		_, err := LoadSession()
		if err == nil {
			t.Error("expected error for invalid json")
		}
	})

	t.Run("valid", func(t *testing.T) {
		data := `{"repo_dir":"repo","queries_dir":"queries","methods":[]}`
		os.WriteFile(filepath.Join(".codist", "sql-session.json"), []byte(data), 0o644)
		sess, err := LoadSession()
		if err != nil {
			t.Fatalf("LoadSession() error: %v", err)
		}
		if sess.RepoDir != "repo" {
			t.Errorf("expected RepoDir 'repo', got %q", sess.RepoDir)
		}
	})
}
