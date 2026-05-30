//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestParseOneField_Plain 테스트
package actix

import "testing"

func TestParseOneField_Plain(t *testing.T) {
	src := []byte(`struct S { name: String }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	fds := collectFieldDecls(root)
	if len(fds) != 1 {
		t.Fatalf("expected 1 field decl, got %d", len(fds))
	}
	f := parseOneField(fds[0], src, nil)
	if f == nil {
		t.Fatal("expected field")
	}
	if f.Name != "name" {
		t.Errorf("name = %q", f.Name)
	}
	if f.Nullable {
		t.Error("expected not nullable")
	}
}
