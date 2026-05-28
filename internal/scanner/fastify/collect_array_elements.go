//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 배열 노드에서 구분자를 제외한 요소 노드만 수집한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func collectArrayElements(arr *sitter.Node) []*sitter.Node {
	var elems []*sitter.Node
	for i := 0; i < int(arr.ChildCount()); i++ {
		child := arr.Child(i)
		t := child.Type()
		if t == "[" || t == "]" || t == "," {
			continue
		}
		elems = append(elems, child)
	}
	return elems
}
