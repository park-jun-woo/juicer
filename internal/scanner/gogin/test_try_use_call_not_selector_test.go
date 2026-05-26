//ff:func feature=scan type=extract control=sequence
//ff:what TestTryUseCall_NotSelector 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestTryUseCall_NotSelector(t *testing.T) {
	call := &ast.CallExpr{Fun: &ast.Ident{Name: "fn"}}
	routers := make(map[string]*routerInfo)
	tryUseCall(call, routers)
	// Should not panic
}
