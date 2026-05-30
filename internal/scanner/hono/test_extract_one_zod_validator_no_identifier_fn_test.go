//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOneZodValidator_NoIdentifierFn 테스트
package hono

import "testing"

func TestExtractOneZodValidator_NoIdentifierFn(t *testing.T) {

	call, src := firstCallExpr(t, `a.b("json", s);`)
	if v := extractOneZodValidator(call, src); v != nil {
		t.Fatalf("expected nil, got %+v", v)
	}
}
