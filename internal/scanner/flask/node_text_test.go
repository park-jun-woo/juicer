//ff:func feature=scan type=test control=sequence topic=flask
//ff:what nodeText 테스트
package flask

import "testing"

func TestNodeText(t *testing.T) {
	b := []byte("foo = 1\n")
	root, err := parsePython(b)
	if err != nil {
		t.Fatal(err)
	}
	ids := findAllByType(root, "identifier")
	if len(ids) == 0 {
		t.Fatal("no identifier")
	}
	if got := nodeText(ids[0], b); got != "foo" {
		t.Fatalf("nodeText = %q, want foo", got)
	}
}
