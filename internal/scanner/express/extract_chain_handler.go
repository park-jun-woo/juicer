//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what 체인 호출의 arguments에서 핸들러명과 미들웨어 목록을 추출한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func extractChainHandlerAndMiddleware(call *sitter.Node, src []byte) (string, []string, *sitter.Node, string, []string) {
	args := findChildByType(call, "arguments")
	if args == nil {
		return "", nil, nil, "public", nil
	}
	argNodes := collectArgNodes(args)
	if len(argNodes) == 0 {
		return "", nil, nil, "public", nil
	}
	lastArg := argNodes[len(argNodes)-1]
	handler := extractHandlerName(lastArg, src)
	var middleware []string
	for i := 0; i < len(argNodes)-1; i++ {
		mw := extractMiddlewareName(argNodes[i], src)
		if mw != "" {
			middleware = append(middleware, mw)
		}
	}
	authLevel, roles := extractAuthFromMiddlewareNodes(argNodes[:len(argNodes)-1], src)
	return handler, middleware, lastArg, authLevel, roles
}
