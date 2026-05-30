//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what exprStmtFor 테스트 헬퍼
package fastapi

import "testing"

func exprStmtFor(t *testing.T, src []byte) (*pydanticField, bool) {
	t.Helper()
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	stmts := findAllByType(root, "expression_statement")
	if len(stmts) == 0 {
		return nil, false
	}
	return extractFieldFromExprStmt(stmts[0], src), true
}
