//ff:func feature=scan type=test control=sequence
//ff:what TestExprString_IndexExprCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestExprString_IndexExprCov(t *testing.T) {
	got := exprString(&ast.IndexExpr{X: &ast.Ident{Name: "arr"}, Index: &ast.Ident{Name: "i"}})
	if got != "arr[i]" {
		t.Fatalf("expected arr[i], got %s", got)
	}
}
