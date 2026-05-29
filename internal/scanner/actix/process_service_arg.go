//ff:func feature=scan type=extract control=selection topic=actix
//ff:what web::scope/web::resource 호출 인자를 분기 처리해 라우트를 수집한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func processServiceArg(callExpr *sitter.Node, src []byte, prefix string, routes *[]builderRoute) {
	switch findCallRoot(callExpr, src) {
	case "web::scope":
		scopePrefix := extractScopePrefix(callExpr, src)
		collectServiceCalls(callExpr, src, joinPath(prefix, scopePrefix), routes)
	case "web::resource":
		resourcePath := extractResourcePath(callExpr, src)
		collectRouteCalls(callExpr, src, joinPath(prefix, resourcePath), routes)
	}
}
