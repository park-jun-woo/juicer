//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what export 구문의 specifier 중 바인딩명(alias 우선, 없으면 name)이 importName과 일치하는 것이 있는지 검사한다
package express

import sitter "github.com/smacker/go-tree-sitter"

// reexportHasBinding — export 구문의 specifier 바인딩명이 importName과 같은 것이 있으면 true.
func reexportHasBinding(stmt *sitter.Node, src []byte, importName string) bool {
	for _, sp := range findAllByType(stmt, "export_specifier") {
		bind := sp.ChildByFieldName("alias")
		if bind == nil {
			bind = sp.ChildByFieldName("name")
		}
		if bind != nil && nodeText(bind, src) == importName {
			return true
		}
	}
	return false
}
