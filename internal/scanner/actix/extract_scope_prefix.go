//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what web::scope("...") 호출의 첫 문자열 인자(스코프 프리픽스)를 추출한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func extractScopePrefix(node *sitter.Node, src []byte) string {
	return firstStringArgOfScopedCall(node, src, "web::scope")
}
