//ff:func feature=scan type=parse control=sequence topic=zod
//ff:what call_expression에서 function 필드 또는 fallback 자식을 찾는다
package zod

import sitter "github.com/smacker/go-tree-sitter"

func resolveFunctionNode(node *sitter.Node) *sitter.Node {
	fn := node.ChildByFieldName("function")
	if fn != nil {
		return fn
	}
	fn = findChildByType(node, "member_expression")
	if fn != nil {
		return fn
	}
	return findChildByType(node, "identifier")
}
