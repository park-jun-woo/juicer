//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what 노드의 첫 type 자식 텍스트를 반환한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// typeText returns the text of the first type child.
func typeText(node *sitter.Node, src []byte) string {
	t := findChildByType(node, "type")
	if t == nil {
		return ""
	}
	return nodeText(t, src)
}
