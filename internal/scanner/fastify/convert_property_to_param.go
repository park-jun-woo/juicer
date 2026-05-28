//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what JSON Schema property pair를 scanner.Param으로 변환한다
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

func convertPropertyToParam(pair *sitter.Node, src []byte) *scanner.Param {
	name := pairKeyName(pair, src)
	val := pairValueNode(pair)
	if name == "" || val == nil {
		return nil
	}
	p := scanner.Param{Name: name, Type: "string"}
	if val.Type() == "object" {
		typStr := extractPairStringOrIdent(val, src, "type")
		if typStr != "" {
			p.Type = mapJSONSchemaType(typStr)
		}
		defStr := extractPairStringOrIdent(val, src, "default")
		if defStr != "" {
			p.Default = defStr
		}
	}
	return &p
}
