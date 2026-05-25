//ff:func feature=hurl type=session control=sequence
//ff:what TestLoadSession_NoFile 테스트
package hurls

import (
	"os"
	"testing"
)

func TestLoadSession_NoFile(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	_, err := LoadSession()
	if err == nil {
		t.Fatal("expected error")
	}
}
