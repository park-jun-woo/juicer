//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what assignment 노드에 Field(...) 호출이 있는지 확인한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// hasFieldCall checks if an assignment has a Field(...) call on the right side.
func hasFieldCall(assign *sitter.Node, src []byte) bool {
	callNode := findChildByType(assign, "call")
	if callNode == nil {
		return false
	}
	funcIdent := findChildByType(callNode, "identifier")
	return funcIdent != nil && nodeText(funcIdent, src) == "Field"
}
