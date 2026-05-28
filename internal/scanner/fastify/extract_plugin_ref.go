//ff:func feature=scan type=extract control=selection topic=fastify
//ff:what register() 인자 노드에서 플러그인 참조를 추출한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func extractPluginRef(node *sitter.Node, src []byte) string {
	switch node.Type() {
	case "identifier":
		return nodeText(node, src)
	case "call_expression":
		return extractCallStringArg(node, src)
	case "arrow_function", "function_expression", "function":
		return "(inline)"
	}
	return ""
}
