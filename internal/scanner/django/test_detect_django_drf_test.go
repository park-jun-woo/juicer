//ff:func feature=scan type=test control=sequence topic=django
//ff:what requirements.txt에 djangorestframework가 있으면 django로 감지한다
package django

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestDetectDjango_WithDRF(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "requirements.txt"), []byte("django==4.2\ndjangorestframework==3.14\n"), 0o644)

	fw := scanner.DetectFramework(dir)
	if fw != "django" {
		t.Errorf("expected django, got %q", fw)
	}
}
