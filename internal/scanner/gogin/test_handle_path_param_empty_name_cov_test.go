//ff:func feature=scan type=test control=sequence
//ff:what TestHandlePathParam_EmptyNameCov 테스트
package gogin

import (
	"go/ast"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestHandlePathParam_EmptyNameCov(t *testing.T) {
	ep := &scanner.Endpoint{}
	handlePathParam(ep, &ast.CallExpr{Args: []ast.Expr{&ast.Ident{Name: "x"}}})
}
