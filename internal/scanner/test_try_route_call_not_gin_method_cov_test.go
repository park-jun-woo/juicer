//ff:func feature=scan type=test control=sequence
//ff:what TestTryRouteCall_NotGinMethodCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestTryRouteCall_NotGinMethodCov(t *testing.T) {
	call := &ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "r"}, Sel: &ast.Ident{Name: "Run"}}}
	_, ok := tryRouteCall(call, nil, "", nil)
	if ok {
		t.Fatal("expected false")
	}
}
