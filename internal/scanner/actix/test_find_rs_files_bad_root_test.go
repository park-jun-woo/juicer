//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestFindRsFiles_BadRoot 테스트
package actix

import (
	"path/filepath"
	"testing"
)

func TestFindRsFiles_BadRoot(t *testing.T) {

	_, err := findRsFiles(filepath.Join(t.TempDir(), "does-not-exist"))
	if err == nil {
		t.Fatal("expected error for non-existent root")
	}
}
