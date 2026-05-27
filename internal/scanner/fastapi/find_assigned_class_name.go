//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 변수에 할당된 클래스 이름을 AST에서 찾는다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// findAssignedClassName finds a pattern like `varName = ClassName()` in the AST
// and returns "ClassName". Returns "" if no such assignment exists.
func findAssignedClassName(root *sitter.Node, varName string, src []byte) string {
	assignments := findAllByType(root, "assignment")
	for _, assign := range assignments {
		left := findChildByType(assign, "identifier")
		if left == nil || nodeText(left, src) != varName {
			continue
		}
		call := findChildByType(assign, "call")
		if call == nil {
			continue
		}
		ident := findChildByType(call, "identifier")
		if ident != nil {
			return nodeText(ident, src)
		}
	}
	return ""
}
