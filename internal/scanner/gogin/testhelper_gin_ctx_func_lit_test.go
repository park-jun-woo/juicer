//ff:func feature=scan type=test control=sequence
//ff:what ginCtxFuncLit 테스트 헬퍼
package gogin

import "go/ast"

func ginCtxFuncLit(name string) *ast.FuncLit {
	return &ast.FuncLit{
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{
					{
						Names: []*ast.Ident{{Name: name}},
						Type: &ast.StarExpr{
							X: &ast.SelectorExpr{
								X:   &ast.Ident{Name: "gin"},
								Sel: &ast.Ident{Name: "Context"},
							},
						},
					},
				},
			},
		},
		Body: &ast.BlockStmt{},
	}
}
