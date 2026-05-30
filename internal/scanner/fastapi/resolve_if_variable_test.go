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

func TestResolveIfVariable_Branches(t *testing.T) {
	src := []byte("PREFIX = \"/v1\"\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}

	// empty -> unchanged
	if got := resolveIfVariable(root, "", src); got != "" {
		t.Errorf("empty: got %q", got)
	}
	// not an identifier (contains dot) -> unchanged
	if got := resolveIfVariable(root, "settings.API", src); got != "settings.API" {
		t.Errorf("non-ident: got %q", got)
	}
	// identifier resolved against assignment
	if got := resolveIfVariable(root, "PREFIX", src); got != "/v1" {
		t.Errorf("resolved: got %q, want /v1", got)
	}
	// identifier with no assignment -> original returned
	if got := resolveIfVariable(root, "UNKNOWN", src); got != "UNKNOWN" {
		t.Errorf("unresolved: got %q", got)
	}
}
