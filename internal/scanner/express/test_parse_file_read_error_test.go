//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestParseFile_ReadError 테스트
package express

import (
	"path/filepath"
	"testing"
)

func TestParseFile_ReadError(t *testing.T) {
	_, err := parseFile(filepath.Join(t.TempDir(), "missing.ts"))
	if err == nil {
		t.Fatal("expected read error")
	}
}
