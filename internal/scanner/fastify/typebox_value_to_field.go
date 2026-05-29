//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what TypeBox 값 call_expression(Optional 언랩 후)을 scanner.Field로 변환한다
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

func typeBoxValueToField(name string, call *sitter.Node, src []byte) scanner.Field {
	f := scanner.Field{Name: name, JSON: name}
	tbName := typeBoxCallName(call, src)
	if scalar := mapTypeBoxType(tbName); scalar != "" {
		f.Type = scalar
		applyTypeBoxOptions(&f, call, src)
		return f
	}
	if tbName == "Array" {
		f.Type = "array"
		applyTypeBoxArray(&f, call, src)
		return f
	}
	if tbName == "Object" {
		f.Type = "object"
		f.Fields = typeBoxObjectToFields(typeBoxFirstArg(call), src)
		return f
	}
	return f
}
