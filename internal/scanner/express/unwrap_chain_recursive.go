//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 체인의 재귀 단계에서 내부 호출을 풀고 현재 HTTP 메서드를 추가한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func unwrapChainRecursive(innerCall, outerCall *sitter.Node, propName string, src []byte, routers map[string]bool) (string, string, []chainMethod) {
	innerPath, routerVar, innerMethods := unwrapChain(innerCall, src, routers)
	if innerPath == "" {
		return "", "", nil
	}
	upperMethod, ok := httpMethods[propName]
	if !ok {
		return innerPath, routerVar, innerMethods
	}
	handler, mw, hNode, authLevel, roles := extractChainHandlerAndMiddleware(outerCall, src)
	joiRefs := chainJoiRefs(outerCall, src)
	zodValidators := chainZodValidators(outerCall, src)
	return innerPath, routerVar, append(innerMethods, chainMethod{method: upperMethod, handler: handler, handlerNode: hNode, middleware: mw, line: int(outerCall.StartPoint().Row) + 1, authLevel: authLevel, roles: roles, joiRefs: joiRefs, zodValidators: zodValidators})
}
