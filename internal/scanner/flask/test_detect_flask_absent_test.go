//ff:func feature=scan type=test control=sequence topic=flask
//ff:what flask가 없는 requirements.txt는 flask로 감지하지 않는다
package flask

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestDetectFlask_WithoutFlask(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "requirements.txt"), []byte("django==4.0\nredis\n"), 0o644)

	fw := scanner.DetectFramework(dir)
	if fw == "flask" {
		t.Errorf("expected non-flask, got %q", fw)
	}
}
