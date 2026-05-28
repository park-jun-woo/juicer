//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what argument_list에서 첫 번째 identifier 인자를 추출한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// firstIdentArg finds the first positional identifier argument in an argument list.
func firstIdentArg(args *sitter.Node, src []byte) string {
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.Type() == "identifier" {
			return nodeText(child, src)
		}
	}
	return ""
}
