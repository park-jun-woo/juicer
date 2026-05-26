//ff:func feature=scan type=test control=sequence
//ff:what TestTryUseCall_NonSelectorCov 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestTryUseCall_NonSelectorCov(t *testing.T) {
	call := &ast.CallExpr{Fun: &ast.Ident{Name: "f"}}
	tryUseCall(call, nil)
}
