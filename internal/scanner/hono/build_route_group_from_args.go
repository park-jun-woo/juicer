//ff:func feature=scan type=extract control=sequence topic=hono
//ff:what .route() 호출의 인자에서 prefix와 subApp 이름을 파싱한다
package hono

import sitter "github.com/smacker/go-tree-sitter"

func buildRouteGroupFromArgs(call *sitter.Node, src []byte, parentVar string) *routeGroup {
	args := findChildByType(call, "arguments")
	if args == nil {
		return nil
	}
	argNodes := collectArgNodes(args)
	if len(argNodes) < 2 {
		return nil
	}
	prefixNode := argNodes[0]
	if prefixNode.Type() != "string" {
		return nil
	}
	subAppNode := argNodes[1]
	if subAppNode.Type() != "identifier" {
		return nil
	}
	return &routeGroup{
		Prefix:     unquoteTS(nodeText(prefixNode, src)),
		ParentVar:  parentVar,
		SubAppName: nodeText(subAppNode, src),
	}
}
