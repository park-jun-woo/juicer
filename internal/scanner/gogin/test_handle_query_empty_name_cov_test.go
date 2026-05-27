//ff:func feature=scan type=test control=sequence
//ff:what TestHandleQuery_EmptyNameCov 테스트
package gogin

import (
	"go/ast"
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandleQuery_EmptyNameCov(t *testing.T) {
	ep := &scanner.Endpoint{}
	handleQuery(ep, &ast.CallExpr{Args: []ast.Expr{&ast.Ident{Name: "x"}}}, "Query")
}
