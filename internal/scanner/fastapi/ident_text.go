//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what 노드의 첫 identifier 자식 텍스트를 반환한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// identText returns the text of the first identifier child.
func identText(node *sitter.Node, src []byte) string {
	id := findChildByType(node, "identifier")
	if id == nil {
		return ""
	}
	return nodeText(id, src)
}
