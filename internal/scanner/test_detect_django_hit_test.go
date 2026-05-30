//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestDetectDjango_Hit 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectDjango_Hit(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "requirements.txt"), []byte("Django==4.2\n"), 0o644)
	if !detectDjango(dir) {
		t.Fatal("expected true")
	}
}
