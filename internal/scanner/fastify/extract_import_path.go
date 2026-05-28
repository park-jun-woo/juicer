//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what import 구문에서 모듈 경로 문자열을 추출한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func extractImportPath(stmt *sitter.Node, src []byte) string {
	strNode := findChildByType(stmt, "string")
	if strNode == nil {
		return ""
	}
	return unquoteTS(nodeText(strNode, src))
}
