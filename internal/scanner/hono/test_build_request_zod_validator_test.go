//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestBuildRequest_ZodValidator 테스트
package hono

import (
	"github.com/park-jun-woo/codistill/internal/scanner/zod"
	"testing"
)

func TestBuildRequest_ZodValidator(t *testing.T) {
	fi := mustParse(t, []byte(`const s = z.object({ name: z.string() });`+"\n"))
	calls := findAllByType(fi.Root, "call_expression")
	if len(calls) == 0 {
		t.Skip("no call_expression")
	}

	r := routeInfo{
		Method: "POST",
		ZodValidators: []zod.ValidatorInfo{
			{Target: "json", SchemaNode: calls[0]},
		},
	}
	ctx := minimalCtx()
	req := buildRequest(r, nil, ctx, fi)
	if req == nil || req.Body == nil || len(req.Body.Fields) == 0 {
		t.Fatalf("expected body from zod validator, got %+v", req)
	}
}
