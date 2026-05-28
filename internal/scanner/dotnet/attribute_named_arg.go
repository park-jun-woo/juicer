//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what 어트리뷰트에서 이름이 지정된 인자 값을 반환한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func attributeNamedArg(attr *sitter.Node, src []byte, name string) string {
	args := findChildByType(attr, "attribute_argument_list")
	if args == nil {
		return ""
	}
	for _, arg := range childrenOfType(args, "attribute_argument") {
		assign := findChildByType(arg, "assignment_expression")
		if assign == nil {
			continue
		}
		idNode := findChildByType(assign, "identifier")
		if idNode == nil || nodeText(idNode, src) != name {
			continue
		}
		strLit := findChildByType(assign, "string_literal")
		if strLit != nil {
			return unquoteCSharp(nodeText(strLit, src))
		}
	}
	return ""
}
