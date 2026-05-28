//ff:func feature=scan type=extract control=sequence topic=spring
//ff:what 노드가 enum 선언인지 확인한다
package spring

import sitter "github.com/smacker/go-tree-sitter"

func isEnumClass(cls *sitter.Node) bool {
	return cls.Type() == "enum_declaration"
}
