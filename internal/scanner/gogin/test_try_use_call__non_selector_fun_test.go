//ff:func feature=scan type=extract control=sequence
//ff:what TestTryUseCall_NonSelectorFun 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestTryUseCall_NonSelectorFun(t *testing.T) {
	call := &ast.CallExpr{Fun: &ast.Ident{Name: "f"}}
	tryUseCall(call, nil)
}
