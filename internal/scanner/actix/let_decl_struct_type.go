//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what let_declaration이 지정 변수의 struct_expression 바인딩이면 타입명을 반환한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func letDeclStructType(decl *sitter.Node, varName string, src []byte) string {
	pattern := findChildByType(decl, "identifier")
	if pattern == nil || nodeText(pattern, src) != varName {
		return ""
	}
	structExpr := findChildByType(decl, "struct_expression")
	if structExpr == nil {
		return ""
	}
	return structExprTypeName(structExpr, src)
}
