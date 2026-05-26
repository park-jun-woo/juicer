//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveBindType_WithAddressOf 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveBindType_WithAddressOf(t *testing.T) {
	call := &ast.CallExpr{
		Args: []ast.Expr{
			&ast.UnaryExpr{Op: token.AND, X: &ast.Ident{Name: "req"}},
		},
	}
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs:  make(map[*ast.Ident]types.Object),
		Uses:  make(map[*ast.Ident]types.Object),
	}
	typeName, fields := resolveBindType(call, info)
	// req is not defined in info, so empty
	_ = typeName
	_ = fields
}
