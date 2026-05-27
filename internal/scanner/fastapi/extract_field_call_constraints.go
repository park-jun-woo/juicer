//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what assignment 노드에서 Field(...) 호출을 찾아 제약 조건을 추출한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// extractFieldCallConstraints finds a Field(...) call in the assignment and extracts its constraints.
func extractFieldCallConstraints(assign *sitter.Node, src []byte, f *pydanticField) {
	callNode := findChildByType(assign, "call")
	if callNode == nil {
		return
	}
	funcIdent := findChildByType(callNode, "identifier")
	if funcIdent == nil || nodeText(funcIdent, src) != "Field" {
		return
	}
	args := findChildByType(callNode, "argument_list")
	if args == nil {
		return
	}
	extractFieldConstraints(args, src, f)
}
