//ff:func feature=hurl type=session control=sequence
//ff:what TestLoadSession_InvalidJSON 테스트
package hurls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadSession_InvalidJSON(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	os.MkdirAll(sessionDirName, 0o755)
	os.WriteFile(filepath.Join(sessionDirName, sessionFileName), []byte("not json"), 0o644)

	_, err := LoadSession()
	if err == nil {
		t.Fatal("expected error")
	}
}
