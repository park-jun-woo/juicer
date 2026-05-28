//ff:func feature=scan type=test control=sequence topic=django
//ff:what fastapi와 django가 동시에 있으면 django로 감지하지 않는다
package django

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestDetectDjango_FastAPIPriority(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "requirements.txt"), []byte("django==4.2\nfastapi==0.100\n"), 0o644)

	fw := scanner.DetectFramework(dir)
	if fw == "django" {
		t.Errorf("expected non-django when fastapi present, got %q", fw)
	}
}
