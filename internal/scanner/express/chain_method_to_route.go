//ff:func feature=scan type=convert control=sequence topic=express
//ff:what 체인 메서드(chainMethod)를 routeInfo로 변환한다
package express

func chainMethodToRoute(m chainMethod, routePath, routerVar string) routeInfo {
	return routeInfo{
		Method:      m.method,
		Path:        routePath,
		Router:      routerVar,
		Handler:     m.handler,
		HandlerNode: m.handlerNode,
		Middleware:  m.middleware,
		Line:        m.line,
		AuthLevel:   m.authLevel,
		Roles:         m.roles,
		JoiRefs:       m.joiRefs,
		ZodValidators: m.zodValidators,
	}
}
