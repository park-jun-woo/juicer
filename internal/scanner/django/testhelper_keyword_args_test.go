//ff:func feature=scan type=test control=sequence topic=django
//ff:what keywordArgs 테스트 헬퍼
package django

import sitter "github.com/smacker/go-tree-sitter"

func keywordArgs(root *sitter.Node) []*sitter.Node {
	var out []*sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if n.Type() == "keyword_argument" {
			out = append(out, n)
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	return out
}
