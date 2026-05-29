//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what 리소스 직속 .to() 인자에서 핸들러를 파싱해 ANY 메서드로 전개한 라우트들을 추가한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func appendToRoute(args *sitter.Node, src []byte, resourcePath string, routes *[]builderRoute) {
	var handler string
	for i := 0; i < int(args.ChildCount()); i++ {
		if name := handlerNameFromArg(args.Child(i), src); name != "" {
			handler = name
			break
		}
	}
	for _, method := range anyMethods() {
		*routes = append(*routes, builderRoute{
			method:  method,
			path:    resourcePath,
			handler: handlerOrAnon(handler, method, resourcePath),
		})
	}
}
