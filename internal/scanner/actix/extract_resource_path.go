//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what web::resource("...") 호출의 첫 문자열 인자(리소스 경로)를 추출한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func extractResourcePath(node *sitter.Node, src []byte) string {
	return firstStringArgOfScopedCall(node, src, "web::resource")
}
