//ff:func feature=scan type=extract control=iteration dimension=1 topic=zod
//ff:what 파일 내 const 선언에서 z.object({...}) Zod 스키마를 수집한다
package zod

import sitter "github.com/smacker/go-tree-sitter"

// CollectSchemas — AST 루트에서 Zod 스키마 맵 수집
func CollectSchemas(root *sitter.Node, src []byte) map[string]*sitter.Node {
	schemas := make(map[string]*sitter.Node)
	for _, decl := range findAllByType(root, "lexical_declaration") {
		CollectSchemaFromDecl(decl, src, schemas)
	}
	return schemas
}
