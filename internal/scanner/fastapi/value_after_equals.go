//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what keyword_argument에서 = 뒤의 값 텍스트를 반환한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// valueAfterEquals returns the text of the node after "=" in a keyword argument.
func valueAfterEquals(kw *sitter.Node, src []byte) string {
	foundEq := false
	for j := 0; j < int(kw.ChildCount()); j++ {
		ch := kw.Child(j)
		if nodeText(ch, src) == "=" {
			foundEq = true
			continue
		}
		if foundEq {
			return nodeText(ch, src)
		}
	}
	return ""
}
