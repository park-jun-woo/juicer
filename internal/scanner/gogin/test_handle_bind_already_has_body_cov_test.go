//ff:func feature=scan type=test control=sequence
//ff:what TestHandleBind_AlreadyHasBodyCov 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandleBind_AlreadyHasBodyCov(t *testing.T) {
	ep := &scanner.Endpoint{Request: &scanner.Request{Body: &scanner.Body{VarName: "existing"}}}
	call := &ast.CallExpr{}
	info := &types.Info{}
	handleBind(ep, call, "ShouldBindJSON", info)
	if ep.Request.Body.VarName != "existing" {
		t.Fatal("expected existing body to be preserved")
	}
}
