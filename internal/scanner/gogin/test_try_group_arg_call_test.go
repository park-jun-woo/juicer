//ff:func feature=scan type=test control=sequence
//ff:what TestTryGroupArgCall_NoGroupArg 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestTryGroupArgCall_NoGroupArg(t *testing.T) {
	call := &ast.CallExpr{
		Fun: &ast.Ident{Name: "doSomething"},
		Args: []ast.Expr{
			&ast.Ident{Name: "x"},
		},
	}
	ctx := &groupArgCtx{
		ginAlias:  "gin",
		routers:   map[string]*routerInfo{},
		fset:      token.NewFileSet(),
		endpoints: []scanner.Endpoint{},
		hmap:      map[int][]ast.Expr{},
		epIndex:   map[struct{ file string; line int }]int{},
	}
	tryGroupArgCall(call, ctx)
}

func TestTryGroupArgCall_GroupArgMatch(t *testing.T) {
	// register(api) where api is a known router -> extractGroupArgPrefix matches,
	// then rescanCalleeWithPrefix runs (empty info -> early return, no panic).
	call := &ast.CallExpr{
		Fun:  &ast.Ident{Name: "register"},
		Args: []ast.Expr{&ast.Ident{Name: "api"}},
	}
	ctx := &groupArgCtx{
		routers: map[string]*routerInfo{"api": {prefix: "/api"}},
		idx:     &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}},
		info:    goginEmptyInfo(),
	}
	tryGroupArgCall(call, ctx)
}

var _ = scanner.Endpoint{}
