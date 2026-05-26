//ff:func feature=scan type=test control=sequence
//ff:what TestTryGroupArgCall_NoGroupArg 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/park-jun-woo/juicer/internal/scanner"
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
