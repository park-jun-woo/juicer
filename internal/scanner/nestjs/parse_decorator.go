//ff:func feature=scan type=parse control=iteration dimension=1 topic=nestjs
//ff:what 데코레이터 노드에서 이름과 인자를 파싱한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// parseDecorator extracts the decorator name and optional string argument.
func parseDecorator(node *sitter.Node, src []byte) decoratorInfo {
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		if child.Type() == "identifier" {
			return decoratorInfo{name: nodeText(child, src)}
		}
		if child.Type() == "call_expression" {
			return parseDecoratorCall(child, src)
		}
	}
	return decoratorInfo{}
}
