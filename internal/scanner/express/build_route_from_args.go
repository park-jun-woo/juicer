//ff:func feature=scan type=extract control=sequence topic=express
//ff:what arguments 노드에서 경로, 핸들러, 미들웨어, validateRequest를 파싱하여 routeInfo를 생성한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func buildRouteFromArgs(args *sitter.Node, src []byte, method string, line int) *routeInfo {
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
	validators := extractZodValidatorsFromArgs(argNodes, src)
	authLevel, roles := extractAuthFromArgs(argNodes, src)
	lastArg := argNodes[len(argNodes)-1]
	return &routeInfo{
		Method:        method,
		Path:          path,
		Handler:       handler,
		HandlerNode:   lastArg,
		Middleware:    middleware,
		Line:          line,
		ZodValidators: validators,
		AuthLevel:     authLevel,
		Roles:         roles,
	}
}

