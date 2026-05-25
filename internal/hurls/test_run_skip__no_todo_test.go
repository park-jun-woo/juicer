//ff:func feature=hurl type=parse control=sequence
//ff:what TestRunSkip_NoTodo 테스트
package hurls

import (
	"os"
	"testing"
)

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
