//ff:func feature=scan type=test topic=fastify control=sequence
//ff:what typeBoxFirstArg 첫 object/call 인자 노드 반환 테스트
package fastify

import "testing"

func TestTypeBoxFirstArg(t *testing.T) {
	// object arg
	fi := mustParse(t, []byte(`Type.String({ default: 'x' })`))
	call := findAllByType(fi.Root, "call_expression")[0]
	arg := typeBoxFirstArg(call)
	if arg == nil || arg.Type() != "object" {
		t.Errorf("object arg: %v", arg)
	}
	// call_expression arg (Array(Type.String()))
	fi2 := mustParse(t, []byte(`Type.Array(Type.String())`))
	calls := findAllByType(fi2.Root, "call_expression")
	outer := calls[0] // outermost Type.Array(...)
	a2 := typeBoxFirstArg(outer)
	if a2 == nil || a2.Type() != "call_expression" {
		t.Errorf("call arg: %v", a2)
	}
	// no qualifying arg
	fi3 := mustParse(t, []byte(`Type.String()`))
	c3 := findAllByType(fi3.Root, "call_expression")[0]
	if typeBoxFirstArg(c3) != nil {
		t.Error("no arg should be nil")
	}
}
