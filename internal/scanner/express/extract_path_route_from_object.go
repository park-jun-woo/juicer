//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 단일 object 노드에서 path(string)와 route(identifier) 프로퍼티를 추출한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func extractPathRouteFromObject(obj *sitter.Node, src []byte) *pathRouteEntry {
	path := extractPairStringValue(obj, src, "path")
	routeVar := extractPairIdentValue(obj, src, "route")
	if path == "" || routeVar == "" {
		return nil
	}
	return &pathRouteEntry{path: path, routeVar: routeVar}
}
