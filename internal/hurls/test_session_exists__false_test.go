//ff:func feature=hurl type=session control=sequence
//ff:what TestSessionExists_False 테스트
package hurls

import (
	"os"
	"testing"
)

func TestSessionExists_False(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	if SessionExists() {
		t.Fatal("expected false")
	}
}
