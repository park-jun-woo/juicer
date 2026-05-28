//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what 클래스 선언 노드에서 타입 파라미터 이름을 추출한다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func extractTypeParams(cls *sitter.Node, src []byte) []string {
	tp := findChildByType(cls, "type_parameters")
	if tp == nil {
		return nil
	}
	var names []string
	children := childrenOfType(tp, "type_parameter")
	for _, child := range children {
		id := findChildByType(child, "type_identifier")
		if id != nil {
			names = append(names, nodeText(id, src))
		}
	}
	return names
}
