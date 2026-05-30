//ff:func feature=scan type=test control=sequence topic=flask
//ff:what parseAliasedImport가 (로컬명, 원본명)을 파싱한다
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

func TestParseAliasedImport_UnexpectedShape(t *testing.T) {
	// pass a node that lacks the expected children -> ("", "")
	src := []byte("x = 1\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	// the whole module node has no dotted_name+identifier pair at direct-child level
	stmts := findAllByType(root, "integer")
	if len(stmts) == 0 {
		t.Skip("no integer node")
	}
	local, orig := parseAliasedImport(stmts[0], src)
	if local != "" || orig != "" {
		t.Fatalf("expected empty pair, got %q %q", local, orig)
	}
}
