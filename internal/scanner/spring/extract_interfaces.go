//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what 클래스 노드에서 구현한 인터페이스 이름 목록을 추출한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func extractInterfaces(cls *sitter.Node, src []byte) []string {
	superInterfaces := findChildByType(cls, "super_interfaces")
	if superInterfaces == nil {
		return nil
	}
	typeList := findChildByType(superInterfaces, "type_list")
	if typeList == nil {
		return nil
	}
	var result []string
	for i := 0; i < int(typeList.ChildCount()); i++ {
		child := typeList.Child(i)
		if child.Type() == "type_identifier" {
			result = append(result, nodeText(child, src))
		}
	}
	return result
}
