//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what scope 호출의 상위 체인을 따라 .service(handler) 핸들러 이름들을 수집한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func collectServiceHandlers(scopeCallNode *sitter.Node, src []byte) []string {
	var handlers []string
	for parent := scopeCallNode.Parent(); parent != nil; parent = parent.Parent() {
		handlers = appendServiceCallHandlers(parent, src, handlers)
	}
	return handlers
}
