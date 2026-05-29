//ff:func feature=scan type=parse control=sequence topic=zod
//ff:what call_expression 노드에서 메서드명과 인자를 추출하고 체인을 재귀 수집한다
package zod

import sitter "github.com/smacker/go-tree-sitter"

func collectChainFromCallExpr(node *sitter.Node, src []byte, methods *[]ChainMethod) {
	fn := resolveFunctionNode(node)
	if fn == nil || fn.Type() != "member_expression" {
		return
	}
	prop := fn.ChildByFieldName("property")
	if prop != nil {
		cm := buildChainMethodFromProp(node, prop, src)
		*methods = append(*methods, cm)
	}
	obj := fn.ChildByFieldName("object")
	collectChainMethodsRecursive(obj, src, methods)
}
