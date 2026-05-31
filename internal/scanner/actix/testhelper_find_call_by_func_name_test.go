//ff:func feature=scan type=test control=sequence topic=actix
//ff:what findCallByFuncName 테스트 헬퍼
package actix

import sitter "github.com/smacker/go-tree-sitter"

// findCallByFuncName returns the first call_expression whose function is a bare
// identifier equal to name (e.g. the `s()` in `c.service(s())`).
func findCallByFuncName(root *sitter.Node, src []byte, name string) *sitter.Node {
	var found *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if found != nil {
			return
		}
		if n.Type() == "call_expression" {
			fn := n.Child(0)
			if fn != nil && fn.Type() == "identifier" && nodeText(fn, src) == name {
				found = n
				return
			}
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	return found
}
