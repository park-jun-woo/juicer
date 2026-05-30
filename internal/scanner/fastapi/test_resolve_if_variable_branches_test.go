//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestResolveIfVariable_Branches 테스트
package fastapi

import "testing"

func TestResolveIfVariable_Branches(t *testing.T) {
	src := []byte("PREFIX = \"/v1\"\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}

	if got := resolveIfVariable(root, "", src); got != "" {
		t.Errorf("empty: got %q", got)
	}

	if got := resolveIfVariable(root, "settings.API", src); got != "settings.API" {
		t.Errorf("non-ident: got %q", got)
	}

	if got := resolveIfVariable(root, "PREFIX", src); got != "/v1" {
		t.Errorf("resolved: got %q, want /v1", got)
	}

	if got := resolveIfVariable(root, "UNKNOWN", src); got != "UNKNOWN" {
		t.Errorf("unresolved: got %q", got)
	}
}
