//ff:func feature=scan type=test control=sequence
//ff:what TestResolveExprType_DefaultCov 테스트
package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestResolveExprType_DefaultCov(t *testing.T) {
	call := &ast.CallExpr{Fun: &ast.Ident{Name: "f"}}
	info := &types.Info{
		Uses:  make(map[*ast.Ident]types.Object),
		Defs:  make(map[*ast.Ident]types.Object),
		Types: map[ast.Expr]types.TypeAndValue{call: {Type: types.Typ[types.String]}},
	}
	resolveExprType(call, info)
}
