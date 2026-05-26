//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what keyword_argument가 지정 이름과 일치하면 값을 반환한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// tryMatchKeyword checks if a keyword_argument matches the given name and returns its value.
func tryMatchKeyword(kw *sitter.Node, name string, src []byte) string {
	keyNode := findChildByType(kw, "identifier")
	if keyNode == nil || nodeText(keyNode, src) != name {
		return ""
	}
	valNode := findChildByType(kw, "string")
	if valNode != nil {
		return unquotePython(nodeText(valNode, src))
	}
	intNode := findChildByType(kw, "integer")
	if intNode != nil {
		return nodeText(intNode, src)
	}
	return valueAfterEquals(kw, src)
}
