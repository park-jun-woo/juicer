//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what JSON Schema 중첩 object 또는 array items의 필드를 적용한다
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

func applyNestedFields(f *scanner.Field, typStr string, propNode *sitter.Node, src []byte) {
	if typStr == "object" {
		f.Fields = jsonSchemaToFields(propNode, src)
		return
	}
	if typStr == "array" {
		applyArrayItems(f, propNode, src)
	}
}
