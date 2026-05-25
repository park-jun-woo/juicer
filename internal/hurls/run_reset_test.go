//ff:func feature=hurl type=parse control=sequence
//ff:what TestRunReset_WithSession 테스트
package hurls

import (
	"os"
	"testing"
)

func TestRunReset_WithSession(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	sess := &Session{Host: "http://localhost", TestsDir: "tests", RepoDir: "repo"}
	SaveSession(sess)

	if err := RunReset(); err != nil {
		t.Fatal(err)
	}
}
