//ff:func feature=scan type=test control=sequence
//ff:what TestHandlePathParam_EmptyNameCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestHandlePathParam_EmptyNameCov(t *testing.T) {
	ep := &Endpoint{}
	handlePathParam(ep, &ast.CallExpr{Args: []ast.Expr{&ast.Ident{Name: "x"}}})
}
