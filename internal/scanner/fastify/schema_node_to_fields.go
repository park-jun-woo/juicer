//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what 스키마 노드를 Field로 변환한다 (TypeBox 변수 참조면 TypeBox 변환, 아니면 JSON Schema 변환)
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

func schemaNodeToFields(node *sitter.Node, src []byte, vars map[string]*sitter.Node) []scanner.Field {
	if tb := resolveTypeBoxRef(node, src, vars); tb != nil {
		return typeBoxObjectToFields(tb, src)
	}
	return jsonSchemaToFields(node, src)
}
