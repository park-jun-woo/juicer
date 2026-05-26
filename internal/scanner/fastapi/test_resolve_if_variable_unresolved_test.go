//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestResolveIfVariable_VariableUnresolved 테스트
package fastapi

import "testing"

func TestResolveIfVariable_VariableUnresolved(t *testing.T) {
	src := []byte("x = 1\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	got := resolveIfVariable(root, "MISSING_VAR", src)
	if got != "MISSING_VAR" {
		t.Fatalf("expected MISSING_VAR, got %q", got)
	}
}
