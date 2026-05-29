//ff:func feature=scan type=extract control=sequence topic=zod
//ff:what Zod validator의 스키마를 해석하여 Field 슬라이스를 반환한다
package zod

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

// ResolveValidatorFields — ValidatorInfo → Field 슬라이스.
// inlineSrc: SchemaNode (인라인 스키마)가 속한 소스 바이트.
// schemas: 스키마 변수명 → AST 노드.
// schemaSrc: 스키마 변수명 → 해당 노드의 소스 바이트.
func ResolveValidatorFields(v ValidatorInfo, schemas map[string]*sitter.Node, inlineSrc []byte, schemaSrc map[string][]byte) []scanner.Field {
	if v.SchemaNode != nil {
		return ParseSchema(v.SchemaNode, inlineSrc)
	}
	if v.SchemaName == "" {
		return nil
	}
	node, ok := schemas[v.SchemaName]
	if !ok {
		return nil
	}
	src, ok := schemaSrc[v.SchemaName]
	if !ok {
		return nil
	}
	return ParseSchema(node, src)
}
