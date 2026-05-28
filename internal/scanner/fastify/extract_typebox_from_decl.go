//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what lexical_declaration에서 TypeBox Type.Object() 변수를 추출한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func extractTypeBoxFromDecl(decl *sitter.Node, fi *fileInfo, vars map[string]*sitter.Node) {
	for _, declarator := range findAllByType(decl, "variable_declarator") {
		extractTypeBoxFromDeclarator(declarator, fi, vars)
	}
}
