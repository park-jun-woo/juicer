//ff:func feature=scan type=test control=sequence
//ff:what TestTryUseCall_NotUseCov 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestTryUseCall_NotUseCov(t *testing.T) {
	call := &ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "Run"}}}
	tryUseCall(call, nil)
}
