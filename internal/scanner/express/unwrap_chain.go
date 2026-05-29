//ff:func feature=scan type=extract control=sequence topic=express
//ff:what call_expression 체인을 재귀적으로 풀어 경로와 HTTP 메서드 목록을 추출한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func unwrapChain(node *sitter.Node, src []byte, routers map[string]bool) (string, string, []chainMethod) {
	if node.Type() != "call_expression" {
		return "", "", nil
	}
	mem := findChildByType(node, "member_expression")
	if mem == nil {
		return "", "", nil
	}
	prop := mem.ChildByFieldName("property")
	if prop == nil {
		return "", "", nil
	}
	propName := nodeText(prop, src)
	obj := findChildByType(mem, "call_expression")
	if obj == nil {
		return "", "", nil
	}
	if isRouteCall(obj, src, routers) {
		path, methods := unwrapChainBase(obj, node, propName, src)
		return path, routerVarOfCall(obj, src), methods
	}
	return unwrapChainRecursive(obj, node, propName, src, routers)
}
