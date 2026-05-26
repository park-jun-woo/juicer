//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 파라미터 노드에서 이름과 타입을 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// paramNameAndType extracts the name and type of a parameter node.
func paramNameAndType(param *sitter.Node, src []byte) (string, string) {
	name := ""
	typ := "string"
	for i := 0; i < int(param.ChildCount()); i++ {
		child := param.Child(i)
		if child.Type() == "identifier" {
			name = nodeText(child, src)
		}
		if child.Type() == "type_annotation" {
			typ = extractTypeAnnotation(child, src)
		}
	}
	return name, typ
}
