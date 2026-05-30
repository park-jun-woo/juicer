//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestFindPyFiles_BadRoot 테스트
package fastapi

import (
	"path/filepath"
	"testing"
)

func TestFindPyFiles_BadRoot(t *testing.T) {
	_, err := findPyFiles(filepath.Join(t.TempDir(), "nope"))
	if err == nil {
		t.Fatal("expected error for non-existent root")
	}
}
