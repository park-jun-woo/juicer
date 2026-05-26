//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractFieldFromExprStmt_PlainExpr plain expression nil 반환 분기 테스트
package fastapi

import "testing"

func TestExtractFieldFromExprStmt_PlainExpr(t *testing.T) {
	// expression_statement that is just a function call (no ident, no type)
	src := []byte("class M:\n    print('hello')\n")
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
	if f != nil {
		t.Fatal("expected nil for plain call expression")
	}
}
