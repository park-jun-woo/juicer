//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what variable_declarator의 초기화 값 노드를 반환한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func findInitValue(declarator *sitter.Node) *sitter.Node {
	for i := 0; i < int(declarator.ChildCount()); i++ {
		child := declarator.Child(i)
		if child.Type() == "call_expression" || child.Type() == "new_expression" {
			return child
		}
	}
	return nil
}
