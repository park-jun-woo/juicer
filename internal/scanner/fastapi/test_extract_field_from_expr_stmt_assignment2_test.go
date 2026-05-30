//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractFieldFromExprStmt_Assignment2 테스트
package fastapi

import "testing"

func TestExtractFieldFromExprStmt_Assignment2(t *testing.T) {
	f, ok := exprStmtFor(t, []byte("class M:\n    age: int = 5\n"))
	if !ok || f == nil || f.name != "age" {
		t.Fatalf("got %+v", f)
	}
}
