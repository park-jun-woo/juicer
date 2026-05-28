//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what arguments 노드의 자식 중 구두점을 제외한 인자 노드만 수집한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func collectArgNodes(args *sitter.Node) []*sitter.Node {
	var result []*sitter.Node
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		t := child.Type()
		if t == "(" || t == ")" || t == "," {
			continue
		}
		result = append(result, child)
	}
	return result
}
