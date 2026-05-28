//ff:func feature=scan type=extract control=selection topic=express
//ff:what 라우트 호출의 마지막 인자에서 핸들러 함수명을 추출한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func extractHandlerName(node *sitter.Node, src []byte) string {
	switch node.Type() {
	case "identifier":
		return nodeText(node, src)
	case "member_expression":
		return nodeText(node, src)
	case "call_expression":
		return extractHandlerFromCall(node, src)
	case "arrow_function", "function":
		return "(anonymous)"
	}
	return ""
}
