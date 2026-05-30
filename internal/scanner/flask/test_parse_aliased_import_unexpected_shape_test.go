//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestParseAliasedImport_UnexpectedShape 테스트
package flask

import "testing"

func TestParseAliasedImport_UnexpectedShape(t *testing.T) {

	src := []byte("x = 1\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}

	stmts := findAllByType(root, "integer")
	if len(stmts) == 0 {
		t.Skip("no integer node")
	}
	local, orig := parseAliasedImport(stmts[0], src)
	if local != "" || orig != "" {
		t.Fatalf("expected empty pair, got %q %q", local, orig)
	}
}
