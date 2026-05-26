//ff:func feature=scan type=test control=sequence
//ff:what walkForGroupArgs 전 분기 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestWalkForGroupArgs(t *testing.T) {
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

	// empty stmts
	walkForGroupArgs(nil, ctx)

	// various statement types
	stmts := []ast.Stmt{
		&ast.ExprStmt{X: &ast.Ident{Name: "x"}},          // non-call ExprStmt
		&ast.BlockStmt{List: []ast.Stmt{}},                  // empty block
		&ast.IfStmt{Body: &ast.BlockStmt{List: []ast.Stmt{}}}, // if with no Init/Else
		&ast.ForStmt{Body: &ast.BlockStmt{List: []ast.Stmt{}}},
		&ast.RangeStmt{Body: &ast.BlockStmt{List: []ast.Stmt{}}},
		&ast.SwitchStmt{Body: &ast.BlockStmt{List: []ast.Stmt{}}},
		&ast.TypeSwitchStmt{Body: &ast.BlockStmt{List: []ast.Stmt{}}},
		&ast.SelectStmt{Body: &ast.BlockStmt{List: []ast.Stmt{}}},
		&ast.CaseClause{Body: []ast.Stmt{}},
		&ast.CommClause{Body: []ast.Stmt{}},
	}
	walkForGroupArgs(stmts, ctx)

	// if with Init and Else
	stmts2 := []ast.Stmt{
		&ast.IfStmt{
			Init: &ast.ExprStmt{X: &ast.Ident{Name: "init"}},
			Body: &ast.BlockStmt{List: []ast.Stmt{}},
			Else: &ast.BlockStmt{List: []ast.Stmt{}},
		},
	}
	walkForGroupArgs(stmts2, ctx)
}
