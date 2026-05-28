//ff:func feature=scan type=extract control=iteration dimension=1 topic=hono
//ff:what 단일 lexical_declaration에서 new Hono() 변수를 찾아 수집한다
package hono

import sitter "github.com/smacker/go-tree-sitter"

func collectHonoVarFromDecl(decl *sitter.Node, fi *fileInfo, vars map[string]bool) {
	for _, declarator := range childrenOfType(decl, "variable_declarator") {
		nameNode := declarator.ChildByFieldName("name")
		if nameNode == nil {
			continue
		}
		value := declarator.ChildByFieldName("value")
		if value == nil {
			continue
		}
		if isNewHonoCall(value, fi.Src) || isNewHonoChain(value, fi.Src) {
			vars[nodeText(nameNode, fi.Src)] = true
		}
	}
}
