//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestResolveIfVariable_Empty 테스트
package fastapi

import "testing"

func TestResolveIfVariable_Empty(t *testing.T) {
	src := []byte("x = 1\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	got := resolveIfVariable(root, "", src)
	if got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
