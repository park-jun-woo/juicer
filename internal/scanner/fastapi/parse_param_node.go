//ff:func feature=scan type=parse control=selection topic=fastapi
//ff:what 파라미터 노드에서 이름, 타입, 기본값, 기본값 호출명을 추출한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// parseParamNode extracts name, type, default value, and default call name from a parameter node.
func parseParamNode(param *sitter.Node, src []byte) (name, typeName, defaultVal, defaultCall string) {
	switch param.Type() {
	case "typed_parameter":
		name = identText(param, src)
		typeName = typeText(param, src)
	case "default_parameter":
		name = identText(param, src)
		defaultVal, defaultCall = extractDefault(param, src)
	case "typed_default_parameter":
		name = identText(param, src)
		typeName = typeText(param, src)
		defaultVal, defaultCall = extractDefault(param, src)
	}
	return
}
