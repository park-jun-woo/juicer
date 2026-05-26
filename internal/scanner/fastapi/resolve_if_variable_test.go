//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestResolveIfVariable_Literal 테스트
package fastapi

import "testing"

func TestResolveIfVariable_Literal(t *testing.T) {
	src := []byte("x = 1\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	got := resolveIfVariable(root, "/api", src)
	if got != "/api" {
		t.Fatalf("expected /api, got %q", got)
	}
}
