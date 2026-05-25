package hurls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindTestFile_Found(t *testing.T) {
	dir := t.TempDir()
	content := "GET {{host}}/api/health\nHTTP 200\n"
	os.WriteFile(filepath.Join(dir, "health.hurl"), []byte(content), 0o644)
	got := findTestFile(dir, "GET", "/api/health")
	if got == "" {
		t.Fatal("expected to find test file")
	}
}

func TestFindTestFile_NotFound(t *testing.T) {
	dir := t.TempDir()
	got := findTestFile(dir, "GET", "/api/missing")
	if got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestFindTestFile_NoDir(t *testing.T) {
	got := findTestFile("/nonexistent/dir", "GET", "/path")
	if got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
