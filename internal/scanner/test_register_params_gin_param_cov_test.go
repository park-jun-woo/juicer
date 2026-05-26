//ff:func feature=scan type=test control=sequence
//ff:what TestRegisterParams_GinParamCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestRegisterParams_GinParamCov(t *testing.T) {
	fn := &ast.FuncDecl{
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{{Name: "r"}},
						Type: &ast.StarExpr{
							X: &ast.SelectorExpr{X: &ast.Ident{Name: "gin"}, Sel: &ast.Ident{Name: "Engine"}},
						},
					},
					{
						Names: []*ast.Ident{{Name: "other"}},
						Type:  &ast.Ident{Name: "int"},
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
