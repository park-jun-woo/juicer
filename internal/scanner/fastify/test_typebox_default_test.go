//ff:func feature=scan type=test topic=fastify control=sequence
//ff:what typeBoxDefault 옵션 object의 default 값 추출 테스트
package fastify

import "testing"

func TestTypeBoxDefault(t *testing.T) {
	fi := mustParse(t, []byte(`Type.String({ default: 'admin' })`))
	call := findAllByType(fi.Root, "call_expression")[0]
	if got := typeBoxDefault(call, fi.Src); got != "admin" {
		t.Errorf("got %q", got)
	}
	// no options object
	fi2 := mustParse(t, []byte(`Type.String()`))
	c2 := findAllByType(fi2.Root, "call_expression")[0]
	if got := typeBoxDefault(c2, fi2.Src); got != "" {
		t.Errorf("no opts: got %q", got)
	}
}
