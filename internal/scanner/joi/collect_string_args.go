//ff:func feature=scan type=extract control=iteration dimension=1 topic=joi
//ff:what arguments 노드에서 문자열/숫자 인자를 수집한다
package joi

import sitter "github.com/smacker/go-tree-sitter"

func collectStringArgs(args *sitter.Node, src []byte) []string {
	var result []string
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		t := child.Type()
		if t == "(" || t == ")" || t == "," {
			continue
		}
		result = append(result, unquoteJoi(nodeText(child, src)))
	}
	return result
}
