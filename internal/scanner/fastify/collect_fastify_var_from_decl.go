//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what 단일 lexical_declaration에서 Fastify() 호출 인스턴스를 수집한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func collectFastifyVarFromDecl(decl *sitter.Node, fi *fileInfo, instances map[string]bool) {
	for _, declarator := range findAllByType(decl, "variable_declarator") {
		nameNode := findChildByType(declarator, "identifier")
		if nameNode == nil {
			continue
		}
		callNode := findInitCallExpr(declarator)
		if callNode == nil {
			continue
		}
		fnNode := findChildByType(callNode, "identifier")
		if fnNode == nil {
			continue
		}
		fnName := nodeText(fnNode, fi.Src)
		if fnName == "Fastify" || fnName == "fastify" {
			instances[nodeText(nameNode, fi.Src)] = true
		}
	}
}
