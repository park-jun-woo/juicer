//ff:func feature=hurl type=session control=sequence
//ff:what TestDeleteSession_NoFile 테스트
package hurls

import (
	"os"
	"testing"
)

func TestDeleteSession_NoFile(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	err := DeleteSession()
	if err == nil {
		t.Fatal("expected error")
	}
}
