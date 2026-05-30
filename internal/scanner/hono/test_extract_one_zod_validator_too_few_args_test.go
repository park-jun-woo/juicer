//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOneZodValidator_TooFewArgs 테스트
package hono

import "testing"

func TestExtractOneZodValidator_TooFewArgs(t *testing.T) {
	call, src := firstCallExpr(t, `zValidator("json");`)
	if v := extractOneZodValidator(call, src); v != nil {
		t.Fatalf("expected nil, got %+v", v)
	}
}
