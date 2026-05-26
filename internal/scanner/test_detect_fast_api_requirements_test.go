//ff:func feature=scan type=test control=sequence
//ff:what TestDetectFastAPI_Requirements requirements.txt 감지 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectFastAPI_Requirements(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "requirements.txt"), []byte("fastapi\n"), 0o644)
	if !detectFastAPI(dir) {
		t.Fatal("expected true")
	}
}
