//ff:func feature=scan type=extract control=selection topic=actix
//ff:what 핸들러 인자 노드에서 핸들러 이름을 해석한다(identifier/scoped/generic; 클로저 등은 빈 문자열)
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// handlerNameFromArg resolves the handler name from a .to(...) argument node.
// identifier -> text; scoped_identifier (e.g. view::about) -> last segment;
// generic_function (e.g. api::search::<T>) -> last segment of its inner
// scoped_identifier. Closures and other anonymous forms have no name and
// return "" so the caller can assign a path-based anonymous id.
func handlerNameFromArg(arg *sitter.Node, src []byte) string {
	switch arg.Type() {
	case "identifier":
		return nodeText(arg, src)
	case "scoped_identifier":
		parts := splitScoped(nodeText(arg, src))
		return parts[len(parts)-1]
	case "generic_function":
		if sid := findChildByType(arg, "scoped_identifier"); sid != nil {
			parts := splitScoped(nodeText(sid, src))
			return parts[len(parts)-1]
		}
		if id := findChildByType(arg, "identifier"); id != nil {
			return nodeText(id, src)
		}
	}
	return ""
}
