//ff:func feature=scan type=extract control=sequence topic=dotnet
//ff:what 인자 노드의 object_creation_expression에서 타입명을 추출한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func argType(arg *sitter.Node, src []byte) (string, bool) {
	obj := findChildByType(arg, "object_creation_expression")
	if obj == nil {
		return "", false
	}
	if gn := findChildByType(obj, "generic_name"); gn != nil {
		return unwrapReturnType(nodeText(gn, src))
	}
	if id := findChildByType(obj, "identifier"); id != nil {
		return nodeText(id, src), false
	}
	return "", false
}
