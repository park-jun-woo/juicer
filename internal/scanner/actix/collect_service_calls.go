//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what scope 노드에 체이닝된 .service() 호출들을 찾아 인자를 처리한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func collectServiceCalls(node *sitter.Node, src []byte, prefix string, routes *[]builderRoute, handlerFuncs map[string]*handlerInfo, visited map[string]bool) {
	findServiceCalls(node, src, func(args *sitter.Node) {
		processServiceCallArgs(args, src, prefix, routes, handlerFuncs, visited)
	})
}
