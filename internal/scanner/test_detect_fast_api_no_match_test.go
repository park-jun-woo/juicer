//ff:func feature=scan type=test control=sequence
//ff:what TestDetectFastAPI_NoMatch 불일치 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectFastAPI_NoMatch(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "requirements.txt"), []byte("flask\n"), 0o644)
	if detectFastAPI(dir) {
		t.Fatal("expected false")
	}
}
