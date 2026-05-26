//ff:func feature=scan type=test control=sequence
//ff:what TestForwardRouterCalls_DepthLimit 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestForwardRouterCalls_DepthLimit(t *testing.T) {
	call := &ast.CallExpr{
		Fun:  &ast.Ident{Name: "forward"},
		Args: []ast.Expr{&ast.Ident{Name: "router"}},
	}
	stmts := []ast.Stmt{
		&ast.ExprStmt{X: call},
	}
	ctx := &groupArgCtx{
		info: &types.Info{Uses: map[*ast.Ident]types.Object{}},
		fset: token.NewFileSet(),
		idx:  &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}, info: map[token.Pos]*types.Info{}},
	}
	parent := &routerInfo{prefix: "/api"}
	forwardRouterCalls(stmts, "router", "/api", parent, nil, ctx, 1)
}
