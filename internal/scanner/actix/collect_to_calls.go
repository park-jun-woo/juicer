//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what resource 노드에 직접 체이닝된 .to() 호출들을 찾아 ANY 라우트를 추가한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// collectToCalls handles method-less resource registration
// (web::resource("...").to(handler)). It walks DOWN the resource's receiver
// chain only, so a .to() nested inside a .route(web::get().to(h)) argument is
// not matched (that case is handled by collectRouteCalls).
func collectToCalls(node *sitter.Node, src []byte, resourcePath string, routes *[]builderRoute) {
	walkMethodChain(node, src, "to", func(args *sitter.Node) {
		appendToRoute(args, src, resourcePath, routes)
	})
}
