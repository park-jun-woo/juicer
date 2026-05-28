//ff:func feature=scan type=extract control=sequence topic=hono
//ff:what 단일 call_expression에서 app.route("/prefix", subApp) 그룹을 추출한다
package hono

import sitter "github.com/smacker/go-tree-sitter"

func extractRouteGroup(call *sitter.Node, src []byte, honoVars map[string]bool) *routeGroup {
	mem := findChildByType(call, "member_expression")
	if mem == nil {
		return nil
	}
	obj := findChildByType(mem, "identifier")
	if obj == nil {
		return nil
	}
	parentVar := nodeText(obj, src)
	if !honoVars[parentVar] {
		return nil
	}
	prop := mem.ChildByFieldName("property")
	if prop == nil {
		return nil
	}
	if nodeText(prop, src) != "route" {
		return nil
	}
	return buildRouteGroupFromArgs(call, src, parentVar)
}
