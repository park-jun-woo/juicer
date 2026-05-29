//ff:func feature=scan type=extract control=selection topic=actix
//ff:what 최상위 노드 하나를 처리해 대기 macro 어트리뷰트/완성 라우트 상태를 갱신한다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func consumeMacroChild(child *sitter.Node, fi *fileInfo, routes, pending []macroRoute) ([]macroRoute, []macroRoute) {
	switch {
	case child.Type() == "attribute_item":
		r := parseMacroAttribute(child, fi.src)
		if r != nil {
			pending = append(pending, *r)
		}
		return routes, pending
	case child.Type() == "function_item" && len(pending) > 0:
		handler := macroHandlerName(child, fi.src)
		return appendMacroRoutes(routes, pending, child, fi, handler), nil
	default:
		return routes, nil
	}
}
