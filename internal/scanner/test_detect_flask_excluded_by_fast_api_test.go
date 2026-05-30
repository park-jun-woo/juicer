//ff:func feature=scan type=test control=sequence
//ff:what TestDetectFlask_ExcludedByFastAPI 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectFlask_ExcludedByFastAPI(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "requirements.txt"), []byte("flask\nfastapi\n"), 0o644)
	if detectFlask(dir) {
		t.Fatal("expected false when fastapi present")
	}
}
