//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what 필드 선언이 static 키워드를 가지는지 확인한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func isStaticField(field *sitter.Node) bool {
	mods := findModifiers(field)
	if mods == nil {
		return false
	}
	for i := 0; i < int(mods.ChildCount()); i++ {
		child := mods.Child(i)
		if child.Type() == "static" {
			return true
		}
	}
	return false
}
