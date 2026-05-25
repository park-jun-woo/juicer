//ff:func feature=sql type=command control=selection
//ff:what TestHandleSQLSubcommand 테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func setupSQLSession(t *testing.T) (string, func()) {
	t.Helper()
	dir := t.TempDir()
	repoDir := filepath.Join(dir, "repo")
	queriesDir := filepath.Join(dir, "queries")
	os.MkdirAll(repoDir, 0o755)
	os.MkdirAll(queriesDir, 0o755)
	sessionDir := filepath.Join(dir, ".huma")
	os.MkdirAll(sessionDir, 0o755)
	sessionJSON := `{"repo_dir":"` + repoDir + `","queries_dir":"` + queriesDir + `","methods":[]}`
	os.WriteFile(filepath.Join(sessionDir, "sql-session.json"), []byte(sessionJSON), 0o644)
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	return dir, func() { os.Chdir(oldWd) }
}

func TestHandleSQLSubcommand_Next(t *testing.T) {
	_, cleanup := setupSQLSession(t)
	defer cleanup()
	got := handleSQLSubcommand([]string{"next"})
	if !got {
		t.Fatal("expected true")
	}
}

func TestHandleSQLSubcommand_Status(t *testing.T) {
	_, cleanup := setupSQLSession(t)
	defer cleanup()
	got := handleSQLSubcommand([]string{"status"})
	if !got {
		t.Fatal("expected true")
	}
}

func TestHandleSQLSubcommand_List(t *testing.T) {
	_, cleanup := setupSQLSession(t)
	defer cleanup()
	got := handleSQLSubcommand([]string{"list"})
	if !got {
		t.Fatal("expected true")
	}
}

func TestHandleSQLSubcommand_Skip(t *testing.T) {
	_, cleanup := setupSQLSession(t)
	defer cleanup()
	got := handleSQLSubcommand([]string{"skip"})
	if !got {
		t.Fatal("expected true")
	}
}

func TestHandleSQLSubcommand_Reset(t *testing.T) {
	_, cleanup := setupSQLSession(t)
	defer cleanup()
	got := handleSQLSubcommand([]string{"reset"})
	if !got {
		t.Fatal("expected true")
	}
}

func TestHandleSQLSubcommand_Unknown(t *testing.T) {
	got := handleSQLSubcommand([]string{"unknown"})
	if got {
		t.Fatal("expected false for unknown subcommand")
	}
}
