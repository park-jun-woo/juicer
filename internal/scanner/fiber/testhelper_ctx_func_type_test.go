//ff:func feature=scan type=test control=sequence
//ff:what ctxFuncType 테스트 헬퍼
package fiber

import "go/ast"

func ctxFuncType(name string) *ast.FuncType {
	return &ast.FuncType{
		Params: &ast.FieldList{
			List: []*ast.Field{
				{
					Names: []*ast.Ident{{Name: name}},
					Type: &ast.StarExpr{
						X: &ast.SelectorExpr{
							X:   &ast.Ident{Name: "fiber"},
							Sel: &ast.Ident{Name: "Ctx"},
						},
					},
				},
			},
		},
	}
}
