//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what TypeBox 객체 pair를 scanner.Param으로 변환한다 (querystring/params용, Optional 언랩)
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

func typeBoxPairToParam(pair *sitter.Node, src []byte) *scanner.Param {
	name := pairKeyName(pair, src)
	val := pairValueNode(pair)
	if name == "" || val == nil || val.Type() != "call_expression" {
		return nil
	}
	if typeBoxCallName(val, src) == "Optional" {
		val = typeBoxFirstArg(val)
		if val == nil || val.Type() != "call_expression" {
			return &scanner.Param{Name: name, Type: "string"}
		}
	}
	p := scanner.Param{Name: name, Type: "string"}
	if scalar := mapTypeBoxType(typeBoxCallName(val, src)); scalar != "" {
		p.Type = scalar
	}
	if def := typeBoxDefault(val, src); def != "" {
		p.Default = def
	}
	return &p
}
