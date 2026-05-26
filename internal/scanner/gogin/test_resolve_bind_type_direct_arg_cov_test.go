//ff:func feature=scan type=test control=sequence
//ff:what TestResolveBindType_DirectArgCov 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestResolveBindType_DirectArgCov(t *testing.T) {
	ident := &ast.Ident{Name: "req"}
	call := &ast.CallExpr{Args: []ast.Expr{ident}}
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	resolveBindType(call, info)
}
