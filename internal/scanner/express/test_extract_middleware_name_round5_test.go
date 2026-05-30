//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExtractMiddlewareName_Round5 테스트
package express

import "testing"

func TestExtractMiddlewareName_Round5(t *testing.T) {

	f := mustParse(t, []byte(`x = authMiddleware;`))
	if got := extractMiddlewareName(rhsExpr(t, f), f.Src); got != "authMiddleware" {
		t.Errorf("identifier: %q", got)
	}

	f2 := mustParse(t, []byte(`x = auth.required;`))
	if got := extractMiddlewareName(rhsExpr(t, f2), f2.Src); got != "auth.required" {
		t.Errorf("member: %q", got)
	}

	f3 := mustParse(t, []byte(`x = requireRole("admin");`))
	if got := extractMiddlewareName(rhsExpr(t, f3), f3.Src); got != "requireRole" {
		t.Errorf("call: %q", got)
	}
}
