//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what argument_list에서 첫 번째 식별자 인자를 반환한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// firstIdentArg returns the first positional identifier argument text.
// For attribute nodes (e.g., items.router), the full dotted text is returned.
func firstIdentArg(args *sitter.Node, src []byte) string {
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.Type() == "identifier" {
			return nodeText(child, src)
		}
		if child.Type() == "attribute" {
			return nodeText(child, src)
		}
	}
	return ""
}
