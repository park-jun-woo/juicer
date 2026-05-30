//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestStructNameMatches 테스트
package actix

import "testing"

func TestStructNameMatches(t *testing.T) {
	src := []byte(`struct User { id: i64 }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	sn := firstStructNode(root)
	if sn == nil {
		t.Fatal("no struct")
	}
	if !structNameMatches(sn, src, "User") {
		t.Error("expected match for User")
	}
	if structNameMatches(sn, src, "Other") {
		t.Error("expected no match for Other")
	}
}
