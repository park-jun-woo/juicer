//ff:func feature=scan type=extract control=sequence topic=django
//ff:what call 노드에서 함수 이름을 추출한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// callFuncName returns the function name of a call node.
func callFuncName(callNode *sitter.Node, src []byte) string {
	if id := findChildByType(callNode, "identifier"); id != nil {
		return nodeText(id, src)
	}
	if attr := findChildByType(callNode, "attribute"); attr != nil {
		return nodeText(attr, src)
	}
	return ""
}
