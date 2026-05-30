//ff:func feature=scan type=test control=sequence topic=actix
//ff:what collectAttrItems 테스트 헬퍼
package actix

import sitter "github.com/smacker/go-tree-sitter"

func collectAttrItems(root *sitter.Node) []*sitter.Node {
	var items []*sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if n.Type() == "attribute_item" {
			items = append(items, n)
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	return items
}
