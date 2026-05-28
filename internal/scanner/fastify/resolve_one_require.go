//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what 단일 variable_declarator에서 require() 경로를 해석하여 매핑에 추가한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func resolveOneRequire(declarator *sitter.Node, src []byte, dir string, imports map[string]string) {
	nameNode := findChildByType(declarator, "identifier")
	if nameNode == nil {
		return
	}
	varName := nodeText(nameNode, src)
	callNode := findInitCallExpr(declarator)
	if callNode == nil {
		return
	}
	fnNode := findChildByType(callNode, "identifier")
	if fnNode == nil || nodeText(fnNode, src) != "require" {
		return
	}
	requirePath := extractCallStringArg(callNode, src)
	if requirePath == "" {
		return
	}
	resolved := resolveRelativePath(dir, requirePath)
	if resolved != "" {
		imports[varName] = resolved
	}
}
