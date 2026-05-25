package hurls

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

	os.MkdirAll(sessionDirName, 0o755)
	data := `{"host":"http://localhost","tests_dir":"tests","repo_dir":"repo","endpoints":[]}`
	os.WriteFile(filepath.Join(sessionDirName, sessionFileName), []byte(data), 0o644)

	sess, err := LoadSession()
	if err != nil {
		t.Fatal(err)
	}
	if sess.Host != "http://localhost" {
		t.Fatalf("unexpected host: %s", sess.Host)
	}
}

func TestLoadSession_NoFile(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	_, err := LoadSession()
	if err == nil {
		t.Fatal("expected error")
	}
}

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
