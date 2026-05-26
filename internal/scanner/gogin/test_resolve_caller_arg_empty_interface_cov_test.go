//ff:func feature=scan type=test control=sequence
//ff:what TestResolveCallerArg_EmptyInterfaceCov 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestResolveCallerArg_EmptyInterfaceCov(t *testing.T) {
	ifc := types.NewInterfaceType(nil, nil)
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue), Uses: make(map[*ast.Ident]types.Object)}
	r := resolveCallerArg(ifc, &ast.Ident{Name: "data"}, info)
	_ = r
}
