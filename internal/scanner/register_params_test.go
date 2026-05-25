package scanner

import (
	"go/ast"
	"testing"
)

func TestRegisterParams_EmptyParams(t *testing.T) {
	fn := &ast.FuncDecl{Type: &ast.FuncType{}}
	routers := map[string]*routerInfo{}
	registerParams(fn, "gin", routers)
	if len(routers) != 0 {
		t.Fatal("expected no routers")
	}
}

func TestRegisterParams_WithGinEngine(t *testing.T) {
	fn := &ast.FuncDecl{
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{{Name: "r"}},
						Type: &ast.StarExpr{
							X: &ast.SelectorExpr{
								X:   &ast.Ident{Name: "gin"},
								Sel: &ast.Ident{Name: "Engine"},
							},
						},
					},
				},
			},
		},
	}
	routers := map[string]*routerInfo{}
	registerParams(fn, "gin", routers)
	if _, ok := routers["r"]; !ok {
		t.Fatal("expected router r")
	}
}
