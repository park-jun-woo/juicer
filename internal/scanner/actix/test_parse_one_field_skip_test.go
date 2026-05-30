//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestParseOneField_Skip 테스트
package actix

import "testing"

func TestParseOneField_Skip(t *testing.T) {
	src := []byte(`struct S { internal: String }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	fds := collectFieldDecls(root)

	if f := parseOneField(fds[0], src, []serdeAttr{{skip: true}}); f != nil {
		t.Fatalf("expected nil for skipped field, got %+v", f)
	}
}
