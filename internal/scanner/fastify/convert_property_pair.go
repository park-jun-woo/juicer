//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what JSON Schema property pair를 scanner.Field로 변환한다
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

func convertPropertyPair(pair *sitter.Node, src []byte, requiredSet map[string]bool) *scanner.Field {
	name := pairKeyName(pair, src)
	val := pairValueNode(pair)
	if name == "" || val == nil {
		return nil
	}
	f := jsonSchemaPropertyToField(name, val, src)
	if requiredSet[name] {
		f.Validate = "required"
	}
	return &f
}
