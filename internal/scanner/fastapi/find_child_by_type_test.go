//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what findChildByType 테스트
package fastapi

import "testing"

func TestFindChildByType(t *testing.T) {
	src := []byte("x = 1\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	stmt := findChildByType(root, "expression_statement")
	if stmt == nil {
		t.Fatal("expected expression_statement")
	}
	none := findChildByType(root, "nonexistent_type")
	if none != nil {
		t.Fatal("expected nil")
	}
}
