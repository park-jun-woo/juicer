//ff:func feature=hurl type=parse control=sequence
//ff:what TestHandlePass 테스트
package hurls

import (
	"os"
	"testing"
)

func TestHandlePass(t *testing.T) {
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
			{ID: "GET /b", Status: "TODO"},
		},
	}
	err := handlePass(sess, 0, "test.hurl", "repo", "tests")
	if err != nil {
		t.Fatal(err)
	}
	if sess.Endpoints[0].Status != "DONE" {
		t.Fatalf("expected DONE, got %s", sess.Endpoints[0].Status)
	}
}
