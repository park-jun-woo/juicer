//ff:func feature=scan type=extract control=sequence topic=hono
//ff:what 단일 call_expression에서 HTTP 메서드, 경로, 핸들러, zValidator를 추출한다
package hono

import sitter "github.com/smacker/go-tree-sitter"

func extractOneRoute(call *sitter.Node, src []byte, honoVars map[string]bool) *routeInfo {
	mem := findChildByType(call, "member_expression")
	if mem == nil {
		return nil
	}
	obj := findChildByType(mem, "identifier")
	if obj == nil {
		return nil
	}
	ownerVar := nodeText(obj, src)
	if !honoVars[ownerVar] {
		return nil
	}
	prop := mem.ChildByFieldName("property")
	if prop == nil {
		return nil
	}
	upperMethod, ok := httpMethods[nodeText(prop, src)]
	if !ok {
		return nil
	}
	args := findChildByType(call, "arguments")
	if args == nil {
		return nil
	}
	argNodes := collectArgNodes(args)
	if len(argNodes) < 1 {
		return nil
	}
	pathNode := argNodes[0]
	if pathNode.Type() != "string" {
		return nil
	}
	path := unquoteTS(nodeText(pathNode, src))
	handler, middleware := extractHandlerAndMiddleware(argNodes, src)
	validators := extractZodValidators(argNodes, src)
	methods := expandAllMethod(upperMethod)
	if len(methods) == 1 {
		return &routeInfo{
			Method:        methods[0],
			Path:          path,
			Handler:       handler,
			OwnerVar:      ownerVar,
			Middleware:    middleware,
			Line:          int(call.StartPoint().Row) + 1,
			ZodValidators: validators,
		}
	}
	// "all" -> first method stored; expansion happens at endpoint building
	return &routeInfo{
		Method:        upperMethod,
		Path:          path,
		Handler:       handler,
		OwnerVar:      ownerVar,
		Middleware:    middleware,
		Line:          int(call.StartPoint().Row) + 1,
		ZodValidators: validators,
	}
}
