package hurls

import (
	"os"
	"testing"
)

func TestRunNext_AllDone(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	sess := &Session{
		Host:      "http://localhost",
		TestsDir:  "tests",
		RepoDir:   "repo",
		Endpoints: []EndpointStatus{},
	}
	SaveSession(sess)

	if err := RunNext("", "", ""); err != nil {
		t.Fatal(err)
	}
}

func TestRunNext_NoSession_MissingFlags(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	err := RunNext("", "", "")
	if err == nil {
		t.Fatal("expected error for missing flags")
	}
}
