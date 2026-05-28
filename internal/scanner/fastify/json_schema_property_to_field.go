//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what JSON Schema property 노드를 scanner.Field로 변환한다
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

func jsonSchemaPropertyToField(name string, propNode *sitter.Node, src []byte) scanner.Field {
	f := scanner.Field{Name: name, JSON: name}
	if propNode.Type() != "object" {
		return f
	}
	typStr := extractPairStringOrIdent(propNode, src, "type")
	f.Type = mapJSONSchemaType(typStr)
	applyFormat(&f, propNode, src)
	applyNumericConstraints(&f, propNode, src)
	applyEnum(&f, propNode, src)
	applyNestedFields(&f, typStr, propNode, src)
	return f
}
