//ff:func feature=scan type=parse control=sequence topic=zod
//ff:what member_expression 노드에서 z.method 패턴을 확인하고 체인을 재귀 수집한다
package zod

import sitter "github.com/smacker/go-tree-sitter"

func collectChainFromMemberExpr(node *sitter.Node, src []byte, methods *[]ChainMethod) {
	obj := node.ChildByFieldName("object")
	prop := node.ChildByFieldName("property")
	if prop == nil || obj == nil {
		return
	}
	if nodeText(obj, src) == "z" {
		*methods = append(*methods, ChainMethod{Name: nodeText(prop, src)})
		return
	}
	collectChainMethodsRecursive(obj, src, methods)
}
