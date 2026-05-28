//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 단일 call_expression에서 HTTP 메서드, 경로, 핸들러, 미들웨어를 추출한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func extractOneRoute(call *sitter.Node, src []byte, routers map[string]bool) *routeInfo {
	mem := findChildByType(call, "member_expression")
	if mem == nil {
		return nil
	}
	obj := findChildByType(mem, "identifier")
	if obj == nil {
		return nil
	}
	if !routers[nodeText(obj, src)] {
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
	return buildRouteFromArgs(args, src, upperMethod, int(call.StartPoint().Row)+1)
}
