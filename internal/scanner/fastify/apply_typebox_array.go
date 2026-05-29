//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what TypeBox Type.Array(T)의 요소 타입을 필드에 적용한다
package fastify

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

func applyTypeBoxArray(f *scanner.Field, call *sitter.Node, src []byte) {
	inner := typeBoxFirstArg(call)
	if inner == nil || inner.Type() != "call_expression" {
		return
	}
	innerName := typeBoxCallName(inner, src)
	if scalar := mapTypeBoxType(innerName); scalar != "" {
		f.Type = scalar + "[]"
		return
	}
	if innerName == "Object" {
		f.Fields = typeBoxObjectToFields(typeBoxFirstArg(inner), src)
	}
}
