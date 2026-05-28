//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what lexical/variable_declaration에서 이름이 일치하는 함수 선언의 본문을 반환한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func matchVariableDecl(node *sitter.Node, src []byte, name string) *sitter.Node {
	if node.Type() != "lexical_declaration" && node.Type() != "variable_declaration" {
		return nil
	}
	for i := 0; i < int(node.ChildCount()); i++ {
		decl := node.Child(i)
		if decl.Type() != "variable_declarator" {
			continue
		}
		nameNode := decl.ChildByFieldName("name")
		if nameNode == nil || nodeText(nameNode, src) != name {
			continue
		}
		value := decl.ChildByFieldName("value")
		if value == nil {
			continue
		}
		return extractFunctionBody(value)
	}
	return nil
}
