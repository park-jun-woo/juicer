//ff:func feature=scan type=test control=sequence
//ff:what TestTryGroupArgCall_GroupArgUnknownReceiver 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestTryGroupArgCall_GroupArgUnknownReceiver(t *testing.T) {
	call := &ast.CallExpr{
		Fun: &ast.Ident{Name: "doSomething"},
		Args: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   &ast.Ident{Name: "unknown"},
					Sel: &ast.Ident{Name: "Group"},
				},
				Args: []ast.Expr{
					&ast.BasicLit{Kind: token.STRING, Value: `"/api"`},
				},
			},
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
