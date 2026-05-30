//ff:func feature=scan type=test control=sequence topic=echo
//ff:what parseExpr 테스트 헬퍼
package echo

import (
	"go/ast"
	"go/parser"
	"testing"
)

func parseExpr(t *testing.T, code string) ast.Expr {
	t.Helper()
	e, err := parser.ParseExpr(code)
	if err != nil {
		t.Fatal(err)
	}
	return e
}
