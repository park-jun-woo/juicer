//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what resource 노드에 체이닝된 .route() 호출들을 찾아 라우트를 추가한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func collectRouteCalls(node *sitter.Node, src []byte, resourcePath string, routes *[]builderRoute) {
	findRouteCalls(node, src, func(args *sitter.Node) {
		appendRouteFromArgs(args, src, resourcePath, routes)
	})
}
