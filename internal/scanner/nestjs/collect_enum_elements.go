//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what 배열 노드에서 string/number 리터럴 원소를 수집한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// collectEnumElements collects string and number literal elements from an array node.
func collectEnumElements(arr *sitter.Node, src []byte) []string {
	var vals []string
	for j := 0; j < int(arr.ChildCount()); j++ {
		elem := arr.Child(j)
		switch elem.Type() {
		case "string", "template_string":
			vals = append(vals, unquoteTS(nodeText(elem, src)))
		case "number":
			vals = append(vals, nodeText(elem, src))
		}
	}
	return vals
}
