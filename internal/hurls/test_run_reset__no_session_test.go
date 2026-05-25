//ff:func feature=hurl type=parse control=sequence
//ff:what TestRunReset_NoSession 테스트
package hurls

import (
	"os"
	"testing"
)

func TestRunReset_NoSession(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	if err := RunReset(); err != nil {
		t.Fatal(err)
	}
}
