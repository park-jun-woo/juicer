//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what tree-sitter 노드의 소스 텍스트를 반환한다
package actix

import sitter "github.com/smacker/go-tree-sitter"

func nodeText(node *sitter.Node, src []byte) string {
	return node.Content(src)
}
