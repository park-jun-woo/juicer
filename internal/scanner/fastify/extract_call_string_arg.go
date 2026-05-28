//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what call_expression의 첫 번째 문자열 인자를 추출한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func extractCallStringArg(call *sitter.Node, src []byte) string {
	args := findChildByType(call, "arguments")
	if args == nil {
		return ""
	}
	strNode := findChildByType(args, "string")
	if strNode != nil {
		return unquoteTS(nodeText(strNode, src))
	}
	return ""
}
