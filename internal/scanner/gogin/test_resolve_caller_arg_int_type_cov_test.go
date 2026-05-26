//ff:func feature=scan type=test control=sequence
//ff:what TestResolveCallerArg_IntTypeCov 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestResolveCallerArg_IntTypeCov(t *testing.T) {
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue), Uses: make(map[*ast.Ident]types.Object)}
	r := resolveCallerArg(types.Typ[types.Int], &ast.Ident{Name: "x"}, info)
	_ = r
}
