//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what 대입문에서 지정 이름의 RHS 값을 추출한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// extractAssignedValue extracts the RHS value from an assignment with the given LHS name.
func extractAssignedValue(stmtNode *sitter.Node, name string, src []byte) string {
	assign := findChildByType(stmtNode, "assignment")
	if assign == nil {
		return ""
	}
	left := findChildByType(assign, "identifier")
	if left == nil || nodeText(left, src) != name {
		return ""
	}
	for i := 0; i < int(assign.ChildCount()); i++ {
		child := assign.Child(i)
		if child.Type() == "identifier" && nodeText(child, src) != name {
			return nodeText(child, src)
		}
		if child.Type() == "attribute" {
			return nodeText(child, src)
		}
	}
	return ""
}
