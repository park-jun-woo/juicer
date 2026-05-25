//ff:func feature=hurl type=parse control=sequence
//ff:what setupHurlTestSession 함수
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
