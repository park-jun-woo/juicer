//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what TypeBox Type.Object() 변수를 추적하여 JSON Schema AST 노드에 매핑한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func extractTypeBoxVars(fi *fileInfo) map[string]*sitter.Node {
	vars := make(map[string]*sitter.Node)
	for _, decl := range findAllByType(fi.Root, "lexical_declaration") {
		extractTypeBoxFromDecl(decl, fi, vars)
	}
	return vars
}
