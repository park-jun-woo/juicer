package hurls

import (
	"os"
	"testing"
)

func setupHurlTestSession(t *testing.T) func() {
	t.Helper()
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	sess := &Session{
		Host:     "http://localhost",
		TestsDir: "tests",
		RepoDir:  "repo",
		Endpoints: []EndpointStatus{
			{ID: "GET /a", Status: "TODO"},
			{ID: "GET /b", Status: "DONE", TestFile: "b.hurl"},
		},
	}
	SaveSession(sess)
	return func() { os.Chdir(oldWd) }
}

func TestRunList_WithSession(t *testing.T) {
	cleanup := setupHurlTestSession(t)
	defer cleanup()
	if err := RunList(); err != nil {
		t.Fatal(err)
	}
}

func TestRunList_NoSession(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)
	if err := RunList(); err != nil {
		t.Fatal(err)
	}
}
