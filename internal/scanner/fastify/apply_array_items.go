//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what JSON Schema array items의 타입/필드를 적용한다
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

func applyArrayItems(f *scanner.Field, propNode *sitter.Node, src []byte) {
	itemsNode := findPairValue(propNode, src, "items")
	if itemsNode == nil || itemsNode.Type() != "object" {
		return
	}
	itemType := extractPairStringOrIdent(itemsNode, src, "type")
	if itemType == "object" {
		f.Fields = jsonSchemaToFields(itemsNode, src)
	} else {
		f.Type = mapJSONSchemaType(itemType) + "[]"
	}
}
