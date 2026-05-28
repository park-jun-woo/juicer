//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what JSON Schema enum 배열을 필드에 적용한다
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

func applyEnum(f *scanner.Field, propNode *sitter.Node, src []byte) {
	enumNode := findPairValue(propNode, src, "enum")
	if enumNode == nil || enumNode.Type() != "array" {
		return
	}
	for _, elem := range collectArrayElements(enumNode) {
		f.Enum = append(f.Enum, unquoteTS(nodeText(elem, src)))
	}
}
