package hurls

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

	os.MkdirAll(sessionDirName, 0o755)
	os.WriteFile(filepath.Join(sessionDirName, sessionFileName), []byte(`{}`), 0o644)

	if err := DeleteSession(); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteSession_NoFile(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	err := DeleteSession()
	if err == nil {
		t.Fatal("expected error")
	}
}
