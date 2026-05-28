//ff:func feature=scan type=extract control=sequence topic=hono
//ff:what call_expression의 두 번째 인자가 number면 그 값을, 아니면 기본값을 반환한다
package hono

import sitter "github.com/smacker/go-tree-sitter"

func extractSecondNumberArg(call *sitter.Node, src []byte, defaultVal string) string {
	args := findChildByType(call, "arguments")
	if args == nil {
		return defaultVal
	}
	argNodes := collectArgNodes(args)
	if len(argNodes) >= 2 && argNodes[1].Type() == "number" {
		return nodeText(argNodes[1], src)
	}
	return defaultVal
}
