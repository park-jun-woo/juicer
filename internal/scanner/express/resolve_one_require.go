//ff:func feature=scan type=extract control=iteration dimension=1 topic=express
//ff:what lexical_declaration에서 require() 호출을 찾아 변수명→파일 경로를 매핑한다
package express

import sitter "github.com/smacker/go-tree-sitter"

func resolveOneRequire(decl *sitter.Node, src []byte, dir string, imports map[string]string, absRoot string, aliases map[string]string) {
	for _, declarator := range findAllByType(decl, "variable_declarator") {
		nameNode := findChildByType(declarator, "identifier")
		if nameNode == nil {
			continue
		}
		varName := nodeText(nameNode, src)
		callNode := findInitValue(declarator)
		if callNode == nil || callNode.Type() != "call_expression" {
			continue
		}
		fnNode := findChildByType(callNode, "identifier")
		if fnNode == nil || nodeText(fnNode, src) != "require" {
			continue
		}
		requirePath := extractRequirePath(callNode, src)
		if requirePath == "" {
			continue
		}
		resolved := resolveRelativePath(dir, requirePath)
		if resolved == "" {
			resolved = resolvePathAlias(absRoot, requirePath, aliases)
		}
		if resolved != "" {
			imports[varName] = resolved
		}
	}
}
