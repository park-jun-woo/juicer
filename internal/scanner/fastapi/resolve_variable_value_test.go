//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what resolveVariableValue 테스트
package fastapi

import "testing"

func TestResolveVariableValue(t *testing.T) {
	src := []byte(`
API_STR = "/api"
OTHER = 42
router = APIRouter(prefix=API_STR)
`)
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}

	got := resolveVariableValue(root, "API_STR", src)
	if got != "/api" {
		t.Fatalf("expected /api, got %q", got)
	}

	got2 := resolveVariableValue(root, "MISSING", src)
	if got2 != "" {
		t.Fatalf("expected empty for MISSING, got %q", got2)
	}

	got3 := resolveVariableValue(root, "OTHER", src)
	if got3 != "" {
		t.Fatalf("expected empty for non-string OTHER, got %q", got3)
	}
}
