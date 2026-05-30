//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what extractFieldFromExprStmt: 어노테이션만 / assignment / nil 분기
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

func TestExtractFieldFromExprStmt_AnnotationOnly(t *testing.T) {
	// name: str  (no default) -> annotation-only branch
	f, ok := exprStmtFor(t, []byte("class M:\n    name: str\n"))
	if !ok || f == nil || f.name != "name" || f.typeName != "str" {
		t.Fatalf("got %+v", f)
	}
}

func TestExtractFieldFromExprStmt_Assignment2(t *testing.T) {
	f, ok := exprStmtFor(t, []byte("class M:\n    age: int = 5\n"))
	if !ok || f == nil || f.name != "age" {
		t.Fatalf("got %+v", f)
	}
}

func TestExtractFieldFromExprStmt_PlainCall(t *testing.T) {
	// a bare call expression statement -> no assignment, no ident+type -> nil
	f, ok := exprStmtFor(t, []byte("class M:\n    do_something()\n"))
	if !ok {
		t.Skip("no expr stmt")
	}
	if f != nil {
		t.Fatalf("expected nil, got %+v", f)
	}
}
