//ff:func feature=ddl type=test control=sequence
//ff:what TestRun_ParseErrorCov 테스트
package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRun_ParseErrorCov(t *testing.T) {
	dir := t.TempDir()
	f := filepath.Join(dir, "001.up.sql")
	os.WriteFile(f, []byte("data"), 0o644)
	os.Chmod(f, 0o000)
	defer os.Chmod(f, 0o644)
	_, err := Run(dir)
	if err == nil {
		t.Fatal("expected error")
	}
}
