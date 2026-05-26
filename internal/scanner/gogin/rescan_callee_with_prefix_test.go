//ff:func feature=scan type=test control=sequence
//ff:what rescanCalleeWithPrefix 전 분기 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestRescanCalleeWithPrefix(t *testing.T) {
	fset := token.NewFileSet()
	ctx := &groupArgCtx{
		ginAlias:  "gin",
		routers:   map[string]*routerInfo{},
		info:      &types.Info{Uses: map[*ast.Ident]types.Object{}},
		fset:      fset,
		idx:       &funcIndex{},
		root:      "/tmp",
		endpoints: nil,
		epIndex:   map[struct{ file string; line int }]int{},
	}

	// call with unresolvable target -> early return
	call := &ast.CallExpr{Fun: &ast.Ident{Name: "unknown"}}
	parent := &routerInfo{prefix: "/api"}
	rescanCalleeWithPrefix(call, 0, "/api", parent, ctx)
}
