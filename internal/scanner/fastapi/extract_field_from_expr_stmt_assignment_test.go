//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractFieldFromExprStmt_Assignment assignment 분기 테스트
package fastapi

import "testing"

func TestExtractFieldFromExprStmt_Assignment(t *testing.T) {
	src := []byte("class M:\n    name: str = 'default'\n")
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	block := findAllByType(root, "block")
	if len(block) == 0 {
		t.Fatal("no block")
	}
	stmts := childrenOfType(block[0], "expression_statement")
	if len(stmts) == 0 {
		t.Fatal("no expression_statement")
	}
	f := extractFieldFromExprStmt(stmts[0], src)
	if f == nil {
		t.Fatal("expected field from assignment")
	}
}
