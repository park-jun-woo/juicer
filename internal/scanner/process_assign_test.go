package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestProcessAssign_GinInit(t *testing.T) {
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.Ident{Name: "r"}},
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   &ast.Ident{Name: "gin"},
					Sel: &ast.Ident{Name: "Default"},
				},
			},
		},
	}
	routers := map[string]*routerInfo{}
	processAssign(stmt, "gin", routers)
	if _, ok := routers["r"]; !ok {
		t.Fatal("expected router r")
	}
}

func TestProcessAssign_Group(t *testing.T) {
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.Ident{Name: "api"}},
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   &ast.Ident{Name: "r"},
					Sel: &ast.Ident{Name: "Group"},
				},
				Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"/api"`}},
			},
		},
	}
	routers := map[string]*routerInfo{"r": {prefix: ""}}
	processAssign(stmt, "gin", routers)
	if _, ok := routers["api"]; !ok {
		t.Fatal("expected router api")
	}
}

func TestProcessAssign_NonCall(t *testing.T) {
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.Ident{Name: "x"}},
		Rhs: []ast.Expr{&ast.Ident{Name: "y"}},
	}
	routers := map[string]*routerInfo{}
	processAssign(stmt, "gin", routers)
	if len(routers) != 0 {
		t.Fatal("expected no routers")
	}
}

func TestProcessAssign_ExtraRhs(t *testing.T) {
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.Ident{Name: "x"}},
		Rhs: []ast.Expr{&ast.Ident{Name: "a"}, &ast.Ident{Name: "b"}},
	}
	routers := map[string]*routerInfo{}
	processAssign(stmt, "gin", routers)
}
