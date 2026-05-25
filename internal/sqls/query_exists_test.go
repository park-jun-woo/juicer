package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestQueryExists_Found(t *testing.T) {
	dir, err := os.MkdirTemp("", "qe-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	os.WriteFile(filepath.Join(dir, "queries.sql"), []byte("-- name: FindUser :one\nSELECT * FROM users"), 0644)
	if !queryExists(dir, "FindUser") {
		t.Fatal("expected true")
	}
}

func TestQueryExists_NotFound(t *testing.T) {
	dir, err := os.MkdirTemp("", "qe-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	if queryExists(dir, "Missing") {
		t.Fatal("expected false")
	}
}

func TestQueryExists_MissingDir(t *testing.T) {
	if queryExists("/nonexistent", "x") {
		t.Fatal("expected false")
	}
}
