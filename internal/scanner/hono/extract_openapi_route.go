//ff:func feature=scan type=extract control=sequence topic=hono
//ff:what app.openapi(createRoute({ method, path }), handler) 호출에서 routeInfo를 추출한다
package hono

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func extractOpenAPIRoute(call *sitter.Node, src []byte, ownerVar string) *routeInfo {
	obj := findCreateRouteObject(call)
	if obj == nil {
		return nil
	}
	method := strings.ToUpper(readObjectStringProp(obj, "method", src))
	path := readObjectStringProp(obj, "path", src)
	if method == "" || path == "" {
		return nil
	}
	handler := ""
	if args := findChildByType(call, "arguments"); args != nil {
		argNodes := collectArgNodes(args)
		if len(argNodes) >= 2 {
			handler = extractHandlerName(argNodes[len(argNodes)-1], src)
		}
	}
	return &routeInfo{
		Method:        method,
		Path:          path,
		Handler:       handler,
		OwnerVar:      ownerVar,
		Line:          int(call.StartPoint().Row) + 1,
		ZodValidators: extractOpenAPIRequest(obj, src),
	}
}
