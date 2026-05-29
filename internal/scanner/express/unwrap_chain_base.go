//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 체인 베이스(.route() 호출)에서 경로를 추출하고 첫 번째 HTTP 메서드를 반환한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func unwrapChainBase(routeCall, outerCall *sitter.Node, propName string, src []byte) (string, []chainMethod) {
	path := extractRouteCallPath(routeCall, src)
	if path == "" {
		return "", nil
	}
	upperMethod, ok := httpMethods[propName]
	if !ok {
		return "", nil
	}
	handler, mw, hNode, authLevel, roles := extractChainHandlerAndMiddleware(outerCall, src)
	joiRefs := chainJoiRefs(outerCall, src)
	zodValidators := chainZodValidators(outerCall, src)
	return path, []chainMethod{{method: upperMethod, handler: handler, handlerNode: hNode, middleware: mw, line: int(outerCall.StartPoint().Row) + 1, authLevel: authLevel, roles: roles, joiRefs: joiRefs, zodValidators: zodValidators}}
}
