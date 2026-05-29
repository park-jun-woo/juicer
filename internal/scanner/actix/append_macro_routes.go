//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what 대기 중인 macro 어트리뷰트들에 핸들러/함수노드/파일을 채워 라우트 목록에 추가한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func appendMacroRoutes(routes, pending []macroRoute, funcNode *sitter.Node, fi *fileInfo, handler string) []macroRoute {
	for _, attr := range pending {
		attr.handler = handler
		attr.funcNode = funcNode
		attr.file = fi
		routes = append(routes, attr)
	}
	return routes
}
