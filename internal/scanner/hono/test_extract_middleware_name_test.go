//ff:func feature=scan type=test topic=hono control=sequence
//ff:what extractMiddlewareName identifier/call/member 표현식에서 미들웨어명 추출 테스트
package hono

import "testing"

func TestExtractMiddlewareName(t *testing.T) {
	// identifier
	fi := mustParse(t, []byte(`authMiddleware`))
	id := findAllByType(fi.Root, "identifier")[0]
	if got := extractMiddlewareName(id, fi.Src); got != "authMiddleware" {
		t.Errorf("identifier: %q", got)
	}
	// call_expression
	fi2 := mustParse(t, []byte(`cors()`))
	call := findAllByType(fi2.Root, "call_expression")[0]
	if got := extractMiddlewareName(call, fi2.Src); got == "" {
		t.Errorf("call: %q", got)
	}
	// member_expression
	fi3 := mustParse(t, []byte(`mw.auth`))
	mem := findAllByType(fi3.Root, "member_expression")[0]
	if got := extractMiddlewareName(mem, fi3.Src); got != "mw.auth" {
		t.Errorf("member: %q", got)
	}
}
