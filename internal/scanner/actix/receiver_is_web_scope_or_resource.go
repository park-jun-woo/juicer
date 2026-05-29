//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what 수신자 체인 헤드가 web::scope/web::resource 인지 판별한다(인자 내부는 무시)
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func receiverIsWebScopeOrResource(node *sitter.Node, src []byte) bool {
	head := chainHeadCallRoot(node, src)
	return head == "web::scope" || head == "web::resource"
}
