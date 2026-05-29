//ff:func feature=scan type=parse control=selection topic=zod
//ff:what AST를 재귀 순회하여 Zod 메서드 체인을 outer→inner 순서로 수집한다
package zod

import sitter "github.com/smacker/go-tree-sitter"

func collectChainMethodsRecursive(node *sitter.Node, src []byte, methods *[]ChainMethod) {
	if node == nil {
		return
	}
	switch node.Type() {
	case "call_expression":
		collectChainFromCallExpr(node, src, methods)
	case "member_expression":
		collectChainFromMemberExpr(node, src, methods)
	}
}
