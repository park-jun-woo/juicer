//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what app.include_router(router) 호출을 찾는다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// findIncludeRouterCalls finds calls like app.include_router(router, prefix="/api/v1").
func findIncludeRouterCalls(root *sitter.Node, src []byte) []includeCall {
	var calls []includeCall
	callNodes := findAllByType(root, "call")
	for _, call := range callNodes {
		inc := tryParseIncludeRouter(call, src)
		if inc != nil {
			calls = append(calls, *inc)
		}
	}
	return calls
}
