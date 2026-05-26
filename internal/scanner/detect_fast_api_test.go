//ff:func feature=scan type=test control=sequence
//ff:what TestDetectFastAPI_RequirementsHit 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectFastAPI_RequirementsHit(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "requirements.txt"), []byte("fastapi==0.100.0\n"), 0o644)
	if !detectFastAPI(dir) {
		t.Fatal("expected true for fastapi in requirements.txt")
	}
}
