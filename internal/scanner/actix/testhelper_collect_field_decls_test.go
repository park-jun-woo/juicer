//ff:func feature=scan type=test control=sequence topic=actix
//ff:what collectFieldDecls 테스트 헬퍼
package actix

import sitter "github.com/smacker/go-tree-sitter"

func collectFieldDecls(root *sitter.Node) []*sitter.Node {
	var out []*sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if n.Type() == "field_declaration" {
			out = append(out, n)
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	return out
}
