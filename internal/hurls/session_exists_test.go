package hurls

import (
	"os"
	"testing"
)

func TestSessionExists_True(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	sess := &Session{Host: "http://localhost"}
	SaveSession(sess)

	if !SessionExists() {
		t.Fatal("expected true")
	}
}

func TestSessionExists_False(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	if SessionExists() {
		t.Fatal("expected false")
	}
}
