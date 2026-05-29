//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what field_expression에서 web::get/post 등의 HTTP 메서드를 추출한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func extractWebMethod(fieldExpr *sitter.Node, src []byte) string {
	for i := 0; i < int(fieldExpr.ChildCount()); i++ {
		child := fieldExpr.Child(i)
		if m := webMethodFromCall(child, src); m != "" {
			return m
		}
	}
	return ""
}
