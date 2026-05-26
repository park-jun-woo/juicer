//ff:func feature=scan type=test control=sequence
//ff:what TestHandleQuery_EmptyNameCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestHandleQuery_EmptyNameCov(t *testing.T) {
	ep := &Endpoint{}
	handleQuery(ep, &ast.CallExpr{Args: []ast.Expr{&ast.Ident{Name: "x"}}}, "Query")
}
