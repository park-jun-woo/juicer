//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what invocation_expression에서 receiver와 method name을 추출한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func extractMethodCall(inv *sitter.Node, src []byte) (string, string) {
	access := findChildByType(inv, "member_access_expression")
	if access == nil {
		return "", ""
	}
	ids := childrenOfType(access, "identifier")
	if len(ids) < 2 {
		return "", ""
	}
	return nodeText(ids[0], src), nodeText(ids[1], src)
}
