//ff:func feature=hurl type=parse control=sequence
//ff:what TestRunSkip_NoSession 테스트
package hurls

import (
	"os"
	"testing"
)

func TestRunSkip_NoSession(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	if err := RunSkip(); err != nil {
		t.Fatal(err)
	}
}
