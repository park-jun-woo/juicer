//ff:func feature=scan type=extract control=sequence topic=express
//ff:what call_expression에서 호출 함수명을 추출한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func extractHandlerFromCall(node *sitter.Node, src []byte) string {
	fn := findChildByType(node, "identifier")
	if fn != nil {
		return nodeText(fn, src)
	}
	mem := findChildByType(node, "member_expression")
	if mem != nil {
		return nodeText(mem, src)
	}
	return ""
}
