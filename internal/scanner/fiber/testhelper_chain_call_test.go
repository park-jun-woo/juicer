//ff:func feature=scan type=test control=sequence
//ff:what chainCall 테스트 헬퍼
package fiber

import "go/ast"

func chainCall(method string, args []ast.Expr) (*ast.CallExpr, *ast.CallExpr) {
	statusCall := &ast.CallExpr{
		Fun:  &ast.SelectorExpr{X: &ast.Ident{Name: "c"}, Sel: &ast.Ident{Name: "Status"}},
		Args: args,
	}
	outerCall := &ast.CallExpr{
		Fun:  &ast.SelectorExpr{X: statusCall, Sel: &ast.Ident{Name: method}},
		Args: nil,
	}
	return statusCall, outerCall
}
