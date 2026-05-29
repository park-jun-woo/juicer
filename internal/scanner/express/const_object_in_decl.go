//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 단일 lexical_declaration에서 이름이 name인 object 리터럴 value 노드를 찾는다
package express

import sitter "github.com/smacker/go-tree-sitter"

func constObjectInDecl(decl *sitter.Node, src []byte, name string) *sitter.Node {
	for _, declarator := range childrenOfType(decl, "variable_declarator") {
		nameNode := declarator.ChildByFieldName("name")
		if nameNode == nil || nodeText(nameNode, src) != name {
			continue
		}
		value := declarator.ChildByFieldName("value")
		if value != nil && value.Type() == "object" {
			return value
		}
	}
	return nil
}
