package hurls

import (
	"os"
	"testing"
)

func TestRunStatus_WithSession(t *testing.T) {
	cleanup := setupHurlTestSession(t)
	defer cleanup()
	if err := RunStatus(); err != nil {
		t.Fatal(err)
	}
}

func TestRunStatus_NoSession(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	if err := RunStatus(); err != nil {
		t.Fatal(err)
	}
}
