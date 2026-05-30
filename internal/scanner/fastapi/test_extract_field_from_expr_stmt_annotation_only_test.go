//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestExtractFieldFromExprStmt_AnnotationOnly 테스트
package fastapi

import "testing"

func TestExtractFieldFromExprStmt_AnnotationOnly(t *testing.T) {

	f, ok := exprStmtFor(t, []byte("class M:\n    name: str\n"))
	if !ok || f == nil || f.name != "name" || f.typeName != "str" {
		t.Fatalf("got %+v", f)
	}
}
