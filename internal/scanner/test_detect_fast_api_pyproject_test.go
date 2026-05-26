//ff:func feature=scan type=test control=sequence
//ff:what TestDetectFastAPI_Pyproject pyproject.toml 감지 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectFastAPI_Pyproject(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "pyproject.toml"), []byte("[project]\ndependencies=[\"fastapi\"]\n"), 0o644)
	if !detectFastAPI(dir) {
		t.Fatal("expected true")
	}
}
