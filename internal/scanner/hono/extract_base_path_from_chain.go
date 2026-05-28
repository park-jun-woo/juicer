//ff:func feature=scan type=extract control=sequence topic=hono
//ff:what call_expression 체인에서 .basePath("/api") 인자를 추출한다
package hono

import sitter "github.com/smacker/go-tree-sitter"

func extractBasePathFromChain(node *sitter.Node, src []byte) string {
	if node.Type() != "call_expression" {
		return ""
	}
	mem := findChildByType(node, "member_expression")
	if mem == nil {
		return ""
	}
	prop := mem.ChildByFieldName("property")
	if prop == nil {
		return ""
	}
	if nodeText(prop, src) != "basePath" {
		return ""
	}
	if !chainContainsNewHono(node, src) {
		return ""
	}
	return extractFirstStringArg(node, src)
}
