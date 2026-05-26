//ff:func feature=scan type=test control=sequence
//ff:what TestTryRouteCall_NonSelector 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestTryRouteCall_NonSelector(t *testing.T) {
	call := &ast.CallExpr{Fun: &ast.Ident{Name: "f"}}
	_, _, ok := tryRouteCall(call, nil, "", nil)
	if ok {
		t.Fatal("expected false")
	}
}

