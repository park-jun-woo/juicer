//ff:func feature=scan type=test control=sequence
//ff:what tryGroupArgCall 전 분기 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestTryGroupArgCall(t *testing.T) {
	fset := token.NewFileSet()
	ctx := &groupArgCtx{
		ginAlias:  "gin",
		routers:   map[string]*routerInfo{},
		fset:      fset,
		idx:       &funcIndex{},
		root:      "/tmp",
		endpoints: nil,
		epIndex:   map[struct{ file string; line int }]int{},
	}

	// call with non-Group arg
	call := &ast.CallExpr{
		Fun:  &ast.Ident{Name: "setup"},
		Args: []ast.Expr{&ast.Ident{Name: "x"}},
	}
	tryGroupArgCall(call, ctx)
}
