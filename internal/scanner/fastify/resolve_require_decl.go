//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what lexical_declaration에서 require() 호출을 찾아 변수명 -> 파일 경로를 매핑한다
package fastify

import sitter "github.com/smacker/go-tree-sitter"

func resolveRequireDecl(decl *sitter.Node, src []byte, dir string, imports map[string]string) {
	for _, declarator := range findAllByType(decl, "variable_declarator") {
		resolveOneRequire(declarator, src, dir, imports)
	}
}
