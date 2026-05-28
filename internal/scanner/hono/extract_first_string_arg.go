//ff:func feature=scan type=extract control=sequence topic=hono
//ff:what call_expression의 첫 번째 문자열 인자를 추출한다
package hono

import sitter "github.com/smacker/go-tree-sitter"

func extractFirstStringArg(callNode *sitter.Node, src []byte) string {
	args := findChildByType(callNode, "arguments")
	if args == nil {
		return ""
	}
	argNodes := collectArgNodes(args)
	if len(argNodes) < 1 {
		return ""
	}
	if argNodes[0].Type() != "string" {
		return ""
	}
	return unquoteTS(nodeText(argNodes[0], src))
}
