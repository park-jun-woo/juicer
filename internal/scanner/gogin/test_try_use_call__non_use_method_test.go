//ff:func feature=scan type=extract control=sequence
//ff:what TestTryUseCall_NonUseMethod 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestTryUseCall_NonUseMethod(t *testing.T) {
	call := &ast.CallExpr{
		Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "GET"}},
	}
	tryUseCall(call, nil)
}
