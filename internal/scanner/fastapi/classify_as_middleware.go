//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what Depends 타입 별칭 또는 인라인 Annotated를 미들웨어로 분류한다
package fastapi

// classifyAsMiddleware adds the Depends function name to ri.middleware,
// resolving from alias map or from inline Annotated text.
func classifyAsMiddleware(typeName string, aliasMap map[string]string, ri *routeInfo) {
	if aliasMap != nil {
		if fn, ok := aliasMap[typeName]; ok {
			ri.middleware = append(ri.middleware, fn)
			return
		}
	}
	fn := extractDependsFromAnnotated(typeName)
	if fn != "" {
		ri.middleware = append(ri.middleware, fn)
	}
}
