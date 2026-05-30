//ff:func feature=scan type=test control=sequence topic=django
//ff:what findChildByType — 직접 자식 중 지정 타입 첫 노드 반환을 검증
package django

import "testing"

func TestFindChildByType(t *testing.T) {
	src := []byte("x = foo()\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	stmt := firstExprStatement(root)
	if stmt == nil {
		t.Fatal("no expression_statement")
	}
	// found: the assignment child exists.
	if findChildByType(stmt, "assignment") == nil {
		t.Error("expected to find assignment child")
	}
	// not found: there is no class_definition child.
	if findChildByType(stmt, "class_definition") != nil {
		t.Error("expected nil for absent type")
	}
}
