//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what string 노드에서 문자열 내용을 추출한다
package laravel

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// extractStringContent extracts the string content from a string node.
func extractStringContent(node *sitter.Node, src []byte) string {
	strNodes := findAllByType(node, "string_content")
	if len(strNodes) > 0 {
		return nodeText(strNodes[0], src)
	}
	strLit := findChildByType(node, "string")
	if strLit != nil {
		return unquotePHP(nodeText(strLit, src))
	}
	return unquotePHP(nodeText(node, src))
}
