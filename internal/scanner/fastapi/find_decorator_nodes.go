//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 데코레이터에서 call과 attribute 노드를 추출한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// findDecoratorNodes extracts call and attribute nodes from a decorator.
func findDecoratorNodes(dec *sitter.Node) (*sitter.Node, *sitter.Node) {
	var callNode, attrNode *sitter.Node
	for i := 0; i < int(dec.ChildCount()); i++ {
		child := dec.Child(i)
		if child.Type() == "call" {
			callNode = child
		}
		if child.Type() == "attribute" {
			attrNode = child
		}
	}
	return callNode, attrNode
}
