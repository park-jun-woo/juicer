//ff:func feature=scan type=extract control=iteration dimension=1 topic=zod
//ff:what 단일 lexical_declaration에서 Zod 스키마 변수를 찾아 수집한다
package zod

import sitter "github.com/smacker/go-tree-sitter"

// CollectSchemaFromDecl — 단일 선언에서 z.xxx() 스키마 수집
func CollectSchemaFromDecl(decl *sitter.Node, src []byte, schemas map[string]*sitter.Node) {
	for _, declarator := range childrenOfType(decl, "variable_declarator") {
		nameNode := declarator.ChildByFieldName("name")
		if nameNode == nil {
			continue
		}
		value := declarator.ChildByFieldName("value")
		if value == nil {
			continue
		}
		if ContainsCall(value, src) {
			schemas[nodeText(nameNode, src)] = value
		}
	}
}
