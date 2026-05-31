//ff:func feature=scan type=parse control=iteration dimension=1 topic=nestjs
//ff:what 배열 노드에서 string/template_string 자식 값을 수집한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// collectArrayStringArgs collects the string/template_string element values of
// an array node (e.g. ['/a', '/b'] -> ["/a", "/b"]).
func collectArrayStringArgs(arr *sitter.Node, src []byte) []string {
	var out []string
	for j := 0; j < int(arr.ChildCount()); j++ {
		el := arr.Child(j)
		if el.Type() == "string" || el.Type() == "template_string" {
			out = append(out, unquoteTS(nodeText(el, src)))
		}
	}
	return out
}
