//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what Depends() 파라미터를 처리하여 미들웨어로 기록한다
package fastapi

// handleDepends processes a Depends() parameter, extracting the dependency
// function name and classifying it as middleware.
func handleDepends(name, defaultVal string, ri *routeInfo) {
	funcName := extractDependsFuncName(defaultVal)
	if funcName == "" {
		return
	}
	ri.middleware = append(ri.middleware, funcName)
}
