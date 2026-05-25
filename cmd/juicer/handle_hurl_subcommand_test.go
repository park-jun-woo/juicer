//ff:func feature=hurl type=command control=selection
//ff:what TestHandleHurlSubcommand 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func setupHurlSession(t *testing.T) (string, func()) {
	t.Helper()
	dir := t.TempDir()
	sessionDir := filepath.Join(dir, ".huma")
	os.MkdirAll(sessionDir, 0o755)
	sessionJSON := `{"host":"http://localhost","tests_dir":"tests","repo_dir":"repo","endpoints":[]}`
	os.WriteFile(filepath.Join(sessionDir, "hurl-session.json"), []byte(sessionJSON), 0o644)
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	return dir, func() { os.Chdir(oldWd) }
}

func TestHandleHurlSubcommand_Next(t *testing.T) {
	_, cleanup := setupHurlSession(t)
	defer cleanup()
	// next with empty session => "All tests complete!"
	handleHurlSubcommand([]string{"next"})
}

func TestHandleHurlSubcommand_Status(t *testing.T) {
	_, cleanup := setupHurlSession(t)
	defer cleanup()
	handleHurlSubcommand([]string{"status"})
}

func TestHandleHurlSubcommand_List(t *testing.T) {
	_, cleanup := setupHurlSession(t)
	defer cleanup()
	handleHurlSubcommand([]string{"list"})
}

func TestHandleHurlSubcommand_Skip(t *testing.T) {
	_, cleanup := setupHurlSession(t)
	defer cleanup()
	handleHurlSubcommand([]string{"skip"})
}

func TestHandleHurlSubcommand_Reset(t *testing.T) {
	_, cleanup := setupHurlSession(t)
	defer cleanup()
	handleHurlSubcommand([]string{"reset"})
}
