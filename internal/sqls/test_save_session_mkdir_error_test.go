//ff:func feature=sql type=session control=sequence
//ff:what TestSaveSession_MkdirError 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSaveSession_MkdirError(t *testing.T) {
	dir := t.TempDir()
	oldWd, _ := os.Getwd()
	// Create a file where .juicer dir would be
	os.WriteFile(filepath.Join(dir, ".juicer"), []byte("not a dir"), 0o644)
	os.Chdir(dir)
	defer os.Chdir(oldWd)

	sess := &Session{}
	err := SaveSession(sess)
	if err == nil {
		t.Error("expected error when .juicer is a file")
	}
}
