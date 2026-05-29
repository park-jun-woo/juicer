//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what .route() 인자에서 method/handler를 파싱해 라우트 목록에 추가한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func appendRouteFromArgs(args *sitter.Node, src []byte, resourcePath string, routes *[]builderRoute) {
	method, handler := parseRouteArg(args, src)
	if method == "" {
		return
	}
	*routes = append(*routes, builderRoute{
		method:  method,
		path:    resourcePath,
		handler: handlerOrAnon(handler, method, resourcePath),
	})
}
