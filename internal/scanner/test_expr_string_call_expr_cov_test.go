//ff:func feature=scan type=test control=sequence
//ff:what TestExprString_CallExprCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestExprString_CallExprCov(t *testing.T) {
	got := exprString(&ast.CallExpr{Fun: &ast.Ident{Name: "f"}})
	if got != "f()" {
		t.Fatalf("expected f(), got %s", got)
	}
}
