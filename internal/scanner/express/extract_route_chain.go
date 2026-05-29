//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what router.route("/:id").get(h).put(h) 체인 패턴에서 라우트를 추출한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func extractRouteChain(call *sitter.Node, src []byte, routers map[string]bool) []routeInfo {
	routePath, routerVar, methods := unwrapChain(call, src, routers)
	if routePath == "" || len(methods) == 0 {
		return nil
	}
	var routes []routeInfo
	for _, m := range methods {
		routes = append(routes, routeInfo{
			Method:      m.method,
			Path:        routePath,
			Router:      routerVar,
			Handler:     m.handler,
			HandlerNode: m.handlerNode,
			Middleware:  m.middleware,
			Line:        m.line,
			AuthLevel:   m.authLevel,
			Roles:       m.roles,
		})
	}
	return routes
}
