//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestResolveModelFields_ImportResolvesButClassMissing 테스트
package fastapi

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"os"
	"path/filepath"
	"testing"
)

func TestResolveModelFields_ImportResolvesButClassMissing(t *testing.T) {
	dir := t.TempDir()

	modelFile := filepath.Join(dir, "models.py")
	os.WriteFile(modelFile, []byte("x = 1\n"), 0o644)
	referrer := filepath.Join(dir, "main.py")

	cache := make(map[string][]scanner.Field)
	req := modelRequest{
		typeName: "Missing",
		referrer: referrer,
		imports:  []importInfo{{name: "Missing", module: ".models"}},
	}
	if got := resolveModelFields(req, cache, map[string]*fileInfo{}); got != nil {
		t.Fatalf("expected nil for missing class, got %v", got)
	}
}
