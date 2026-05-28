//ff:func feature=scan type=extract control=sequence topic=express
//ff:what .route() 호출의 문자열 인자에서 경로를 추출한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func extractRouteCallPath(node *sitter.Node, src []byte) string {
	args := findChildByType(node, "arguments")
	if args == nil {
		return ""
	}
	strNode := findChildByType(args, "string")
	if strNode == nil {
		return ""
	}
	return unquoteTS(nodeText(strNode, src))
}
