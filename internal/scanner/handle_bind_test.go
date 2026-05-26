//ff:func feature=scan type=test control=sequence
//ff:what TestHandleBind_NoArgs 테스트
package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestHandleBind_NoArgs(t *testing.T) {
	ep := &Endpoint{}
	call := &ast.CallExpr{}
	info := &types.Info{}
	handleBind(ep, call, "ShouldBindJSON", info)
	if ep.Request == nil || ep.Request.Body == nil {
		t.Fatal("expected body to be set")
	}
}
