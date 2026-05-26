//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what argument_list에서 첫 번째 문자열 인자를 추출한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// firstStringArg finds the first positional string argument.
func firstStringArg(args *sitter.Node, src []byte) string {
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.Type() == "string" {
			return unquotePython(nodeText(child, src))
		}
	}
	return ""
}
