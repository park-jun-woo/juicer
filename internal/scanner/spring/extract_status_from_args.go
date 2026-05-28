//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what 어노테이션 인자에서 HTTP 상태 필드를 추출한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func extractStatusFromArgs(args *sitter.Node, src []byte) string {
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.Type() == "field_access" || child.Type() == "identifier" {
			return nodeText(child, src)
		}
	}
	return ""
}
