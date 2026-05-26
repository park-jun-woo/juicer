//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what tree-sitter 노드의 소스 텍스트를 반환한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// nodeText returns the source text of a tree-sitter node.
func nodeText(node *sitter.Node, src []byte) string {
	return node.Content(src)
}
