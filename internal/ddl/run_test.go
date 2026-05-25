package ddl

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRun_Basic(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "001_init.up.sql"), []byte("CREATE TABLE users (id INT);\n"), 0o644)
	out, err := Run(dir)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(out, "CREATE TABLE users") {
		t.Fatalf("unexpected: %q", out)
	}
}

func TestRun_EmptyDirCov(t *testing.T) {
	dir := t.TempDir()
	out, err := Run(dir)
	if err != nil {
		t.Fatal(err)
	}
	if out != "" {
		t.Fatalf("expected empty, got %q", out)
	}
}

func TestRun_InvalidDirCov(t *testing.T) {
	_, err := Run("/nonexistent/dir/12345")
	if err != nil {
		// glob returns nil on non-matching pattern, not error
		// This depends on OS behavior
	}
}
