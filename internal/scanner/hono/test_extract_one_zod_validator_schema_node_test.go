//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestExtractOneZodValidator_SchemaNode 테스트
package hono

import "testing"

func TestExtractOneZodValidator_SchemaNode(t *testing.T) {
	call, src := firstCallExpr(t, `zValidator("query", z.object({ a: z.string() }));`)
	v := extractOneZodValidator(call, src)
	if v == nil || v.Target != "query" || v.SchemaName != "" || v.SchemaNode == nil {
		t.Fatalf("got %+v", v)
	}
}
