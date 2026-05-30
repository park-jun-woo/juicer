//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestFindTSFiles_BadRoot 테스트
package express

import (
	"path/filepath"
	"testing"
)

func TestFindTSFiles_BadRoot(t *testing.T) {
	_, err := findTSFiles(filepath.Join(t.TempDir(), "nope"))
	if err == nil {
		t.Fatal("expected error for non-existent root")
	}
}
