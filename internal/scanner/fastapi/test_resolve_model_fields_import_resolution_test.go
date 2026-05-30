//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestResolveModelFields_ImportResolution 테스트
package fastapi

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"os"
	"path/filepath"
	"testing"
)

func TestResolveModelFields_ImportResolution(t *testing.T) {
	dir := t.TempDir()
	modelFile := filepath.Join(dir, "models.py")
	os.WriteFile(modelFile, []byte("class Item(BaseModel):\n    title: str\n"), 0o644)
	referrer := filepath.Join(dir, "main.py")

	cache := make(map[string][]scanner.Field)
	req := modelRequest{
		typeName: "Item",
		referrer: referrer,
		imports:  []importInfo{{name: "Item", module: ".models"}},
	}

	got := resolveModelFields(req, cache, map[string]*fileInfo{})
	if len(got) != 1 || got[0].Name != "title" {
		t.Fatalf("import resolution failed: %v", got)
	}
}
