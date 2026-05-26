//ff:func feature=scan type=test control=sequence
//ff:what TestHandleBind_AlreadyHasBodyCov 테스트
package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestHandleBind_AlreadyHasBodyCov(t *testing.T) {
	ep := &Endpoint{Request: &Request{Body: &Body{VarName: "existing"}}}
	call := &ast.CallExpr{}
	info := &types.Info{}
	handleBind(ep, call, "ShouldBindJSON", info)
	if ep.Request.Body.VarName != "existing" {
		t.Fatal("expected existing body to be preserved")
	}
}
