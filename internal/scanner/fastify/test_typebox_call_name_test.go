//ff:func feature=scan type=test topic=fastify control=sequence
//ff:what typeBoxCallName Type.X() 호출의 메서드명 추출 테스트
package fastify

import "testing"

func TestTypeBoxCallName(t *testing.T) {
	fi := mustParse(t, []byte(`Type.String()`))
	call := findAllByType(fi.Root, "call_expression")[0]
	if got := typeBoxCallName(call, fi.Src); got != "String" {
		t.Errorf("got %q", got)
	}
	// non-Type object
	fi2 := mustParse(t, []byte(`Other.String()`))
	call2 := findAllByType(fi2.Root, "call_expression")[0]
	if got := typeBoxCallName(call2, fi2.Src); got != "" {
		t.Errorf("non-Type: got %q", got)
	}
	// not a call_expression
	if got := typeBoxCallName(fi.Root, fi.Src); got != "" {
		t.Errorf("non-call: got %q", got)
	}
}
