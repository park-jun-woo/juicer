//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what pair 노드에서 문자열 값을 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// findPairStringValue returns the unquoted string value from a pair node.
func findPairStringValue(pair *sitter.Node, src []byte) string {
	for i := 0; i < int(pair.ChildCount()); i++ {
		child := pair.Child(i)
		switch child.Type() {
		case "string", "template_string":
			return unquoteTS(nodeText(child, src))
		}
	}
	return ""
}
