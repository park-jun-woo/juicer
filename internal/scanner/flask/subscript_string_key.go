//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what subscript 노드의 문자열 인덱스 키를 반환한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// subscriptStringKey returns the string literal used as a subscript index,
// e.g. "username" for request.form['username']. Returns "" if not a string key.
func subscriptStringKey(node *sitter.Node, src []byte) string {
	for i := 0; i < int(node.ChildCount()); i++ {
		child := node.Child(i)
		if child.Type() == "string" {
			return unquotePython(nodeText(child, src))
		}
	}
	return ""
}
