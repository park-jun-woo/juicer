//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestParseOneField_OptionNullable 테스트
package actix

import "testing"

func TestParseOneField_OptionNullable(t *testing.T) {
	src := []byte(`struct S { bio: Option<String> }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	fds := collectFieldDecls(root)
	f := parseOneField(fds[0], src, nil)
	if f == nil {
		t.Fatal("expected field")
	}
	if !f.Nullable {
		t.Error("expected nullable for Option<...>")
	}
}
