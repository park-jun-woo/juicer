package hurls

import (
	"os"
	"testing"
)

func TestRunSkip_WithTodo(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	sess := &Session{
		Host:     "http://localhost",
		TestsDir: "tests",
		RepoDir:  "repo",
		Endpoints: []EndpointStatus{
			{ID: "GET /a", Status: "TODO"},
		},
	}
	SaveSession(sess)

	if err := RunSkip(); err != nil {
		t.Fatal(err)
	}
}

func TestRunSkip_NoTodo(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	sess := &Session{
		Host:      "http://localhost",
		TestsDir:  "tests",
		RepoDir:   "repo",
		Endpoints: []EndpointStatus{{ID: "GET /a", Status: "DONE"}},
	}
	SaveSession(sess)

	if err := RunSkip(); err != nil {
		t.Fatal(err)
	}
}

func TestRunSkip_NoSession(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	if err := RunSkip(); err != nil {
		t.Fatal(err)
	}
}
