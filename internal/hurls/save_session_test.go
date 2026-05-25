//ff:func feature=hurl type=session control=sequence
//ff:what TestSaveSession 테스트
package hurls

import (
	"os"
	"testing"
)

func TestSaveSession(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	sess := &Session{Host: "http://localhost", TestsDir: "tests", RepoDir: "repo"}
	if err := SaveSession(sess); err != nil {
		t.Fatal(err)
	}
	if !SessionExists() {
		t.Fatal("expected session to exist after save")
	}
}
