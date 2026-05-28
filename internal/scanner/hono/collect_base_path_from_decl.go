//ff:func feature=scan type=extract control=iteration dimension=1 topic=hono
//ff:what 단일 lexical_declaration에서 basePath 패턴을 찾아 매핑에 추가한다
package hono

import sitter "github.com/smacker/go-tree-sitter"

func collectBasePathFromDecl(decl *sitter.Node, fi *fileInfo, basePaths map[string]string) {
	for _, declarator := range childrenOfType(decl, "variable_declarator") {
		nameNode := declarator.ChildByFieldName("name")
		if nameNode == nil {
			continue
		}
		value := declarator.ChildByFieldName("value")
		if value == nil {
			continue
		}
		bp := extractBasePathFromChain(value, fi.Src)
		if bp != "" {
			basePaths[nodeText(nameNode, fi.Src)] = bp
		}
	}
}
