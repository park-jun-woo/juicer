//ff:func feature=hurl type=session control=sequence
//ff:what TestLoadSession 테스트
package hurls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadSession(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	os.MkdirAll(sessionDirName, 0o755)
	data := `{"host":"http://localhost","tests_dir":"tests","repo_dir":"repo","endpoints":[]}`
	os.WriteFile(filepath.Join(sessionDirName, sessionFileName), []byte(data), 0o644)

	sess, err := LoadSession()
	if err != nil {
		t.Fatal(err)
	}
	if sess.Host != "http://localhost" {
		t.Fatalf("unexpected host: %s", sess.Host)
	}
}
