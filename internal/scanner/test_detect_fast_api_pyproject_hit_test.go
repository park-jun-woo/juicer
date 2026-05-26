//ff:func feature=scan type=test control=sequence
//ff:what TestDetectFastAPI_PyprojectHit 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectFastAPI_PyprojectHit(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "pyproject.toml"), []byte("[tool.poetry]\nfastapi = \"^0.100\"\n"), 0o644)
	if !detectFastAPI(dir) {
		t.Fatal("expected true for fastapi in pyproject.toml")
	}
}
