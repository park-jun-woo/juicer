//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestResolveIfVariable_VariableResolved 테스트
package fastapi

import "testing"

func TestResolveIfVariable_VariableResolved(t *testing.T) {
	src := []byte("API_STR = \"/api\"\nrouter = APIRouter(prefix=API_STR)\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	got := resolveIfVariable(root, "API_STR", src)
	if got != "/api" {
		t.Fatalf("expected /api, got %q", got)
	}
}
