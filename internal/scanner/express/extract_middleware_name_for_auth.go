//ff:func feature=scan type=extract control=selection topic=express
//ff:what AST 노드에서 미들웨어 함수명을 추출한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func extractMiddlewareNameForAuth(node *sitter.Node, src []byte) (string, bool) {
	switch node.Type() {
	case "identifier":
		return nodeText(node, src), false
	case "member_expression":
		return nodeText(node, src), false
	case "call_expression":
		fn := findChildByType(node, "identifier")
		if fn != nil {
			return nodeText(fn, src), true
		}
		mem := findChildByType(node, "member_expression")
		if mem != nil {
			return nodeText(mem, src), true
		}
		return "", true
	}
	return "", false
}
