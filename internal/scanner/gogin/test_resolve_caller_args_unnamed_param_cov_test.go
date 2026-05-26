//ff:func feature=scan type=test control=sequence
//ff:what TestResolveCallerArgs_UnnamedParamCov 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestResolveCallerArgs_UnnamedParamCov(t *testing.T) {
	fn := &ast.FuncDecl{
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{Type: &ast.Ident{Name: "int"}},
				},
			},
		},
	}
	call := &ast.CallExpr{Args: []ast.Expr{&ast.Ident{Name: "x"}}}
	callerInfo := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue), Uses: make(map[*ast.Ident]types.Object)}
	calleeInfo := &types.Info{Defs: make(map[*ast.Ident]types.Object), Types: make(map[ast.Expr]types.TypeAndValue)}
	resolveCallerArgs(fn, call, callerInfo, calleeInfo)
}
