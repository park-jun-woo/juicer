//ff:func feature=scan type=test control=sequence topic=flask
//ff:what findChildByType 테스트
package flask

import "testing"

func TestFindChildByType(t *testing.T) {
	root, err := parsePython([]byte("x = 1\n"))
	if err != nil {
		t.Fatal(err)
	}
	if got := findChildByType(root, "expression_statement"); got == nil {
		t.Fatal("expected expression_statement child")
	}
	if got := findChildByType(root, "function_definition"); got != nil {
		t.Fatalf("expected nil, got %s", got.Type())
	}
}
