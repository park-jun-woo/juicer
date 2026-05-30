//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestParseFile_ReadError 테스트
package actix

import (
	"path/filepath"
	"testing"
)

func TestParseFile_ReadError(t *testing.T) {
	dir := t.TempDir()
	_, err := parseFile(dir, filepath.Join(dir, "missing.rs"))
	if err == nil {
		t.Fatal("expected read error for missing file")
	}
}
