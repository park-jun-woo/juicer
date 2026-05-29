//ff:func feature=scan type=extract control=iteration dimension=1 topic=zod
//ff:what arguments 노드에서 괄호와 쉼표를 제외한 인자 노드 목록을 반환한다
package zod

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
