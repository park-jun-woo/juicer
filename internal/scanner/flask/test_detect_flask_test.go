//ff:func feature=scan type=test control=sequence topic=flask
//ff:what requirements.txt에 flask가 있으면 flask로 감지한다
package flask

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestDetectFlask_WithFlask(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "requirements.txt"), []byte("flask==3.0.0\nredis\n"), 0o644)

	fw := scanner.DetectFramework(dir)
	if fw != "flask" {
		t.Errorf("expected flask, got %q", fw)
	}
}
