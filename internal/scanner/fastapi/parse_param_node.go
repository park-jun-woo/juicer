//ff:func feature=scan type=parse control=selection topic=fastapi
//ff:what 파라미터 노드에서 이름, 타입, 기본값, 기본값 호출명을 추출한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// parseParamNode extracts name, type, default value, default call name, and isNone from a parameter node.
// isNone is true when the default value is Python's None literal.
func parseParamNode(param *sitter.Node, src []byte) (name, typeName, defaultVal, defaultCall string, isNone bool) {
	switch param.Type() {
	case "typed_parameter":
		name = identText(param, src)
		typeName = typeText(param, src)
	case "default_parameter":
		name = identText(param, src)
		defaultVal, defaultCall, isNone = extractDefault(param, src)
	case "typed_default_parameter":
		name = identText(param, src)
		typeName = typeText(param, src)
		defaultVal, defaultCall, isNone = extractDefault(param, src)
	}
	return
}
