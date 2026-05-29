//ff:func feature=scan type=extract control=iteration dimension=1 topic=actix
//ff:what 함수 파라미터 노드에서 타입 노드를 찾는다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func findParamType(param *sitter.Node) *sitter.Node {
	for i := 0; i < int(param.ChildCount()); i++ {
		child := param.Child(i)
		switch child.Type() {
		case "generic_type", "type_identifier", "scoped_type_identifier",
			"reference_type", "primitive_type":
			return child
		}
	}
	return nil
}
