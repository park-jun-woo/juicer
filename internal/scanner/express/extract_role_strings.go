//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what call_expression의 인자에서 역할 문자열을 추출한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func extractRoleStrings(callNode *sitter.Node, src []byte) []string {
	args := findChildByType(callNode, "arguments")
	if args == nil {
		return nil
	}
	var roles []string
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.Type() == "string" {
			roles = append(roles, unquoteTS(nodeText(child, src)))
		}
	}
	return roles
}
