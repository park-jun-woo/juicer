//ff:func feature=scan type=test control=sequence
//ff:what TestResolveBindType_WithUnaryArgCov 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestResolveBindType_WithUnaryArgCov(t *testing.T) {
	ident := &ast.Ident{Name: "req"}
	call := &ast.CallExpr{Args: []ast.Expr{
		&ast.UnaryExpr{Op: 17, X: ident}, // token.AND = 17
	}}
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	resolveBindType(call, info)
}
