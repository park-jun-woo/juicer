//ff:func feature=scan type=test control=sequence topic=actix
//ff:what findCallByFuncSuffix 테스트 헬퍼
package actix

import sitter "github.com/smacker/go-tree-sitter"

// findCallByFuncSuffix returns the first call_expression whose function text
// ends with the given suffix (e.g. ".to" / ".get").
func findCallByFuncSuffix(root *sitter.Node, src []byte, suffix string) *sitter.Node {
	var found *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if found != nil {
			return
		}
		if n.Type() == "call_expression" {
			fn := n.ChildByFieldName("function")
			if fn != nil {
				txt := nodeText(fn, src)
				if len(txt) >= len(suffix) && txt[len(txt)-len(suffix):] == suffix {
					found = n
					return
				}
			}
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	return found
}
