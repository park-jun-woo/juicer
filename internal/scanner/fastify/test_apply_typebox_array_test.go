//ff:func feature=scan type=test topic=fastify control=sequence
//ff:what applyTypeBoxArray 스칼라[]/객체 배열 요소 타입 반영 테스트
package fastify

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestApplyTypeBoxArray(t *testing.T) {
	// Array(String) -> string[]
	fi := mustParse(t, []byte(`Type.Array(Type.String())`))
	call := findAllByType(fi.Root, "call_expression")[0]
	var f scanner.Field
	applyTypeBoxArray(&f, call, fi.Src)
	if f.Type != "string[]" {
		t.Errorf("scalar array: %+v", f)
	}

	// Array(Object({...})) -> Fields set
	fi2 := mustParse(t, []byte(`Type.Array(Type.Object({ id: Type.Integer() }))`))
	c2 := findAllByType(fi2.Root, "call_expression")[0]
	var o scanner.Field
	applyTypeBoxArray(&o, c2, fi2.Src)
	if len(o.Fields) != 1 {
		t.Errorf("object array fields: %+v", o)
	}
}
