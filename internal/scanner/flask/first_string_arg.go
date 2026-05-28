//ff:func feature=scan type=extract control=iteration dimension=1 topic=flask
//ff:what argument_list에서 첫 번째 문자열 인자를 추출한다
package flask

import sitter "github.com/smacker/go-tree-sitter"

// firstStringArg finds the first positional string argument in an argument list.
func firstStringArg(args *sitter.Node, src []byte) string {
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.Type() == "string" {
			return unquotePython(nodeText(child, src))
		}
	}
	return ""
}
