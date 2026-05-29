//ff:func feature=scan type=extract control=selection topic=dotnet
//ff:what 반환 타입 노드를 매칭하여 endpointInfo에 할당한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func matchReturnType(child *sitter.Node, src []byte, ep *endpointInfo) bool {
	switch child.Type() {
	case "predefined_type":
		if nodeText(child, src) == "void" {
			return true
		}
	case "identifier":
		if actionResultInterfaces[nodeText(child, src)] {
			return true
		}
	case "generic_name":
		raw := nodeText(child, src)
		typeName, isArray := unwrapReturnType(raw)
		if typeName != "" && !actionResultInterfaces[typeName] {
			ep.returnType = typeName
			ep.returnIsArray = isArray
		}
		return true
	}
	return false
}
