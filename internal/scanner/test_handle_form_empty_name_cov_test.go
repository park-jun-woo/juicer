//ff:func feature=scan type=test control=sequence
//ff:what TestHandleForm_EmptyNameCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestHandleForm_EmptyNameCov(t *testing.T) {
	ep := &Endpoint{}
	handleForm(ep, &ast.CallExpr{Args: []ast.Expr{&ast.Ident{Name: "x"}}})
}
