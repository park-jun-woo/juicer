//ff:func feature=scan type=test control=sequence
//ff:what TestResolveCallerArgs_WithParamsCov 테스트
package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestResolveCallerArgs_WithParamsCov(t *testing.T) {
	paramName := &ast.Ident{Name: "code"}
	fn := &ast.FuncDecl{
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{Names: []*ast.Ident{paramName}, Type: &ast.Ident{Name: "int"}},
				},
			},
		},
	}
	call := &ast.CallExpr{Args: []ast.Expr{&ast.Ident{Name: "x"}}}
	callerInfo := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Uses:  make(map[*ast.Ident]types.Object),
	}
	calleeInfo := &types.Info{
		Defs:  map[*ast.Ident]types.Object{paramName: types.NewVar(0, nil, "code", types.Typ[types.Int])},
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	resolveCallerArgs(fn, call, callerInfo, calleeInfo)
}
