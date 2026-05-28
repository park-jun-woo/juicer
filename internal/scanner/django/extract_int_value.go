//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what keyword_argument 노드에서 정수 값을 추출한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// extractIntValue extracts an integer value from a keyword argument node.
func extractIntValue(kwNode *sitter.Node, src []byte) *int {
	intNode := findChildByType(kwNode, "integer")
	if intNode == nil {
		return nil
	}
	text := nodeText(intNode, src)
	v := 0
	for _, c := range text {
		if c >= '0' && c <= '9' {
			v = v*10 + int(c-'0')
		}
	}
	return &v
}
