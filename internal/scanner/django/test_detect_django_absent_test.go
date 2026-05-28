//ff:func feature=scan type=test control=sequence topic=django
//ff:what django가 없는 requirements.txt는 django로 감지하지 않는다
package django

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestDetectDjango_WithoutDjango(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "requirements.txt"), []byte("flask==3.0\nredis\n"), 0o644)

	fw := scanner.DetectFramework(dir)
	if fw == "django" {
		t.Errorf("expected non-django, got %q", fw)
	}
}
