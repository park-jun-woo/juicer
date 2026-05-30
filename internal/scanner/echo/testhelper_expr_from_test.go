//ff:func feature=scan type=test control=sequence topic=echo
//ff:what exprFrom 테스트 헬퍼
package echo

import (
	"go/ast"
	"testing"
)

func exprFrom(t *testing.T, src string) ast.Expr {
	t.Helper()
	return parseExpr(t, src)
}
