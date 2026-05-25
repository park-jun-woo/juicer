//ff:func feature=sql type=session control=sequence
//ff:what TestDeleteSession 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDeleteSession(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	t.Run("no session", func(t *testing.T) {
		err := DeleteSession()
		if err != nil {
			t.Errorf("DeleteSession() error: %v", err)
		}
	})

	t.Run("with session", func(t *testing.T) {
		os.MkdirAll(".huma", 0o755)
		os.WriteFile(filepath.Join(".huma", "sql-session.json"), []byte("{}"), 0o644)
		err := DeleteSession()
		if err != nil {
			t.Fatalf("DeleteSession() error: %v", err)
		}
		if _, err := os.Stat(".huma"); !os.IsNotExist(err) {
			t.Error("expected .huma to be deleted")
		}
	})
}
