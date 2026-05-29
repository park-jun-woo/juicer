//ff:func feature=scan type=extract control=sequence topic=hono
//ff:what app.openapi(createRoute({...})) 호출에서 createRoute 인자의 object 리터럴 노드를 찾는다
package hono

import sitter "github.com/smacker/go-tree-sitter"

func findCreateRouteObject(call *sitter.Node) *sitter.Node {
	args := findChildByType(call, "arguments")
	if args == nil {
		return nil
	}
	argNodes := collectArgNodes(args)
	if len(argNodes) < 1 {
		return nil
	}
	first := argNodes[0]
	if first.Type() != "call_expression" {
		return nil
	}
	inner := findChildByType(first, "arguments")
	if inner == nil {
		return nil
	}
	return findChildByType(inner, "object")
}
