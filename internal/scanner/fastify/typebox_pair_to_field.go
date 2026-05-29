//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what TypeBox 객체 pair를 scanner.Field로 변환한다 (Optional 언랩 + required 판정)
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

func typeBoxPairToField(pair *sitter.Node, src []byte) *scanner.Field {
	name := pairKeyName(pair, src)
	val := pairValueNode(pair)
	if name == "" || val == nil || val.Type() != "call_expression" {
		return nil
	}
	required := true
	if typeBoxCallName(val, src) == "Optional" {
		required = false
		val = typeBoxFirstArg(val)
		if val == nil || val.Type() != "call_expression" {
			return &scanner.Field{Name: name, JSON: name}
		}
	}
	f := typeBoxValueToField(name, val, src)
	if required {
		f.Validate = "required"
	}
	return &f
}
