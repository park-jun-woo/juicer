//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what resolveModelFields 테스트
package fastapi

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestResolveModelFields(t *testing.T) {
	fields := []pydanticField{{name: "email", typeName: "str"}}
	fi := &fileInfo{absPath: "/a.py", models: map[string][]pydanticField{"User": fields}}
	globalModels := map[string]*fileInfo{"User": fi}
	cache := make(map[string][]scanner.Field)

	req := modelRequest{typeName: "User"}
	result := resolveModelFields(req, cache, globalModels)
	if len(result) != 1 {
		t.Fatalf("expected 1 field, got %d", len(result))
	}

	// cache hit
	result2 := resolveModelFields(req, cache, globalModels)
	if len(result2) != 1 {
		t.Fatalf("cache: expected 1 field, got %d", len(result2))
	}

	// not found
	req2 := modelRequest{typeName: "Unknown"}
	result3 := resolveModelFields(req2, cache, globalModels)
	if result3 != nil {
		t.Fatalf("expected nil for unknown, got %v", result3)
	}
}

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
	// globalModels empty -> goes to import resolution -> extractPydanticModel
	got := resolveModelFields(req, cache, map[string]*fileInfo{})
	if len(got) != 1 || got[0].Name != "title" {
		t.Fatalf("import resolution failed: %v", got)
	}
}

func TestResolveModelFields_GlobalEntryMissingModel(t *testing.T) {
	// globalModels has the key but fi.models lacks typeName -> falls through
	// to import resolution which also fails -> nil
	fi := &fileInfo{absPath: "/a.py", models: map[string][]pydanticField{}}
	globalModels := map[string]*fileInfo{"Ghost": fi}
	cache := make(map[string][]scanner.Field)
	req := modelRequest{typeName: "Ghost", referrer: "/x/main.py"}
	if got := resolveModelFields(req, cache, globalModels); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}

func TestResolveModelFields_ImportResolvesButClassMissing(t *testing.T) {
	dir := t.TempDir()
	// file exists and import resolves, but does not define the class -> nil
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

