//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOneZodValidator_SchemaName 테스트
package hono

import "testing"

func TestExtractOneZodValidator_SchemaName(t *testing.T) {
	call, src := firstCallExpr(t, `zValidator("json", createUserSchema);`)
	v := extractOneZodValidator(call, src)
	if v == nil || v.Target != "json" || v.SchemaName != "createUserSchema" {
		t.Fatalf("got %+v", v)
	}
}
