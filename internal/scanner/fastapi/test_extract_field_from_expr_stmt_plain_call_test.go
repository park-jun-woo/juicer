//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractFieldFromExprStmt_PlainCall 테스트
package fastapi

import "testing"

func TestExtractFieldFromExprStmt_PlainCall(t *testing.T) {

	f, ok := exprStmtFor(t, []byte("class M:\n    do_something()\n"))
	if !ok {
		t.Skip("no expr stmt")
	}
	if f != nil {
		t.Fatalf("expected nil, got %+v", f)
	}
}
