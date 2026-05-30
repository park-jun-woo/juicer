//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestParseAliasedImport 테스트
package flask

import "testing"

func TestParseAliasedImport(t *testing.T) {
	src := []byte("from .auth import auth as auth_blueprint\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	nodes := findAllByType(root, "aliased_import")
	if len(nodes) != 1 {
		t.Fatalf("expected 1 aliased_import, got %d", len(nodes))
	}
	local, orig := parseAliasedImport(nodes[0], src)
	if local != "auth_blueprint" {
		t.Errorf("expected local auth_blueprint, got %q", local)
	}
	if orig != "auth" {
		t.Errorf("expected orig auth, got %q", orig)
	}
}
