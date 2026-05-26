//ff:func feature=scan type=extract control=selection topic=fastapi
//ff:what 노드가 타입 또는 기본값 파라미터인지 확인한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// isParamNode returns true if the node is a typed or default parameter.
func isParamNode(node *sitter.Node) bool {
	switch node.Type() {
	case "typed_parameter", "default_parameter", "typed_default_parameter":
		return true
	}
	return false
}
