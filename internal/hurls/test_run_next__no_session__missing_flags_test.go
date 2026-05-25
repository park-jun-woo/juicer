//ff:func feature=hurl type=parse control=sequence
//ff:what TestRunNext_NoSession_MissingFlags 테스트
package hurls

import (
	"os"
	"testing"
)

func TestRunNext_NoSession_MissingFlags(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	err := RunNext("", "", "")
	if err == nil {
		t.Fatal("expected error for missing flags")
	}
}
