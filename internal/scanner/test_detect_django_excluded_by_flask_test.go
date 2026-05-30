//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestDetectDjango_ExcludedByFlask 테스트
package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectDjango_ExcludedByFlask(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "requirements.txt"), []byte("django\nflask\n"), 0o644)
	if detectDjango(dir) {
		t.Fatal("expected false when flask is present")
	}
}
