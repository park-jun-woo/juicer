//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what argument_list에서 위치 인자(non-keyword) 노드를 추출한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// positionalArgs extracts positional (non-keyword) arguments from an argument_list node.
func positionalArgs(args *sitter.Node) []*sitter.Node {
	var result []*sitter.Node
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		t := child.Type()
		if t == "(" || t == ")" || t == "," || t == "keyword_argument" {
			continue
		}
		result = append(result, child)
	}
	return result
}
