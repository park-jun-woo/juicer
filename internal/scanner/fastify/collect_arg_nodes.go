//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what arguments 노드에서 구분자를 제외한 인자 노드만 수집한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func collectArgNodes(args *sitter.Node) []*sitter.Node {
	var nodes []*sitter.Node
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		t := child.Type()
		if t == "," || t == "(" || t == ")" {
			continue
		}
		nodes = append(nodes, child)
	}
	return nodes
}
