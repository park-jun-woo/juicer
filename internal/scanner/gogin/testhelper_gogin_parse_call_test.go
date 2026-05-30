//ff:func feature=scan type=test control=sequence
//ff:what goginParseCall 테스트 헬퍼
package gogin

import (
	"go/ast"
	"go/parser"
	"testing"
)

func goginParseCall(t *testing.T, expr string) *ast.CallExpr {
	t.Helper()
	e, err := parser.ParseExpr(expr)
	if err != nil {
		t.Fatal(err)
	}
	return e.(*ast.CallExpr)
}
