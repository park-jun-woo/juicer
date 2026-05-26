//ff:func feature=scan type=extract control=iteration dimension=1 topic=nestjs
//ff:what arguments 노드에서 첫 번째 문자열 인자를 추출한다
package nestjs

import sitter "github.com/smacker/go-tree-sitter"

// firstStringArg finds the first string argument in an arguments node.
func firstStringArg(args *sitter.Node, src []byte) (string, bool) {
	for i := 0; i < int(args.ChildCount()); i++ {
		arg := args.Child(i)
		if arg.Type() == "string" || arg.Type() == "template_string" {
			return unquoteTS(nodeText(arg, src)), true
		}
	}
	return "", false
}
